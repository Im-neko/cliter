package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	ot "github.com/dghubble/oauth1/twitter"
	"github.com/im-neko/cliter/api/oauth"
	"github.com/im-neko/cliter/logger"
	pb "github.com/im-neko/cliter/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTweetServiceServer
	pb.UnimplementedOAuthServiceServer
}

var config oauth1.Config

func init() {
	consumerKey := os.Getenv("CK")
	consumerSecret := os.Getenv("CS")
	if consumerKey == "" || consumerSecret == "" {
		logger.Error("Required environment variable missing.")
	}
	config = oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    "oob",
		Endpoint:       ot.AuthorizeEndpoint,
	}
	return
}

func main() {
	address := flag.String("address", ":50051", "address to listen on")
	lis, err := net.Listen("tcp", *address)
	if err != nil {
		logger.Error("Failed to listen: %v", err)
		return
	}

	log.Println("Listening on localhost:50051")

	s := grpc.NewServer()
	pb.RegisterTweetServiceServer(s, &server{})
	pb.RegisterOAuthServiceServer(s, &server{})
	healthpb.RegisterHealthServer(s, &healthServer{})
	for key, srv := range s.GetServiceInfo() {
		logger.Info("Service: ", key, srv.Metadata.(string))
		for mkey, mmeta := range srv.Methods {
			logger.Info("Methods: ", mkey, mmeta.Name)
		}
	}
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve: %v", err)
	}

}

func (s *server) SendTweet(ctx context.Context, in *pb.SendTweetRequest) (*pb.SendTweetResponse, error) {
	logger.Info("Client from: %s\n", in.GetTweetContent())
	token := oauth1.NewToken(in.GetAccessToken(), in.GetAccessSecret())
	httpClient := config.Client(oauth1.NoContext, token)
	cl := twitter.NewClient(httpClient)
	tweet, tweetRes, err := cl.Statuses.Update(in.GetTweetContent(), nil)
	logger.Info("tweet: %v\n", tweet)
	logger.Info("res: %v", tweetRes)
	res := &pb.SendTweetResponse{}
	if err != nil {
		logger.Error("failed to tweet: %v", err)
		res.IsSuccess = false
		return res, err
	}
	res.IsSuccess = bool(tweetRes.StatusCode == 200)
	return res, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	logger.Info("Login Request recived\n")
	rt, url, _ := oauth.LoginImpl(&config)
	res := &pb.LoginResponse{
		RequestToken: rt,
		AuthUrl:      url,
		ErrorStr:     "",
	}
	return res, nil
}

func (s *server) ReceivePIN(ctx context.Context, in *pb.ReceivePINRequest) (*pb.ReceivePINResponse, error) {
	logger.Info("ReceivePIN Request recived\n")
	ai, _ := oauth.ReceivePinImpl(&config, in.GetRequestToken(), in.GetVerifier())
	res := &pb.ReceivePINResponse{
		AccessToken:  ai.AccessToken,
		AccessSecret: ai.AccessSecret,
		ErrorStr:     "",
	}
	return res, nil
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	logger.Info("Echo Request recived\n")
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("Unable to get hostname %v", err)
	}
	if hostname != "" {
		grpc.SendHeader(ctx, metadata.Pairs("hostname", hostname))
	}
	res := &pb.EchoResponse{
		Message: in.GetMessage(),
	}
	return res, nil
}

type healthServer struct{}

// Check is used for health checks
func (s *healthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch is not implemented
func (s *healthServer) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}
