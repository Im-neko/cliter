package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	ot "github.com/dghubble/oauth1/twitter"
	"github.com/im-neko/cliter/api/oauth"
	pb "github.com/im-neko/cliter/proto"
	"google.golang.org/grpc"
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
		log.Fatal("Required environment variable missing.")
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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}
	log.Println("Listening on localhost:50051")

	s := grpc.NewServer()
	pb.RegisterTweetServiceServer(s, &server{})
	pb.RegisterOAuthServiceServer(s, &server{})
	for key, srv := range s.GetServiceInfo() {
		log.Println("Service: ", key, srv.Metadata.(string))
		for mkey, mmeta := range srv.Methods {
			log.Println("Methods: ", mkey, mmeta.Name)
		}
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) SendTweet(ctx context.Context, in *pb.SendTweetRequest) (*pb.SendTweetResponse, error) {
	log.Printf("Client from: %s\n", in.GetTweetContent())
	token := oauth1.NewToken(in.GetAccessToken(), in.GetAccessSecret())
	httpClient := config.Client(oauth1.NoContext, token)
	cl := twitter.NewClient(httpClient)
	tweet, tweetRes, err := cl.Statuses.Update(in.GetTweetContent(), nil)
	log.Printf("tweet: %v\n", tweet)
	log.Printf("res: %v", tweetRes)
	res := &pb.SendTweetResponse{}
	if err != nil {
		log.Printf("failed to tweet: %v", err)
		res.IsSuccess = false
		return res, err
	}
	res.IsSuccess = bool(tweetRes.StatusCode == 200)
	return res, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login Request recived\n")
	rt, url, _ := oauth.LoginImpl(&config)
	res := &pb.LoginResponse{
		RequestToken: rt,
		AuthUrl:      url,
		ErrorStr:     "",
	}
	return res, nil
}

func (s *server) ReceivePIN(ctx context.Context, in *pb.ReceivePINRequest) (*pb.ReceivePINResponse, error) {
	log.Printf("ReceivePIN Request recived\n")
	ai, _ := oauth.ReceivePinImpl(&config, in.GetRequestToken(), in.GetVerifier())
	res := &pb.ReceivePINResponse{
		AccessToken:  ai.AccessToken,
		AccessSecret: ai.AccessSecret,
		ErrorStr:     "",
	}
	return res, nil
}
