package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/im-neko/cliter/cmd/oauth"

	pb "github.com/im-neko/cliter/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// AccessFileName is AccessFileName
	// TODO: Ecnrypt
	AccessFileName = "secrets.txt"
)

var (
	twitterClient pb.TweetServiceClient
	ctx           context.Context
	cancel        context.CancelFunc
	rootCmd       = &cobra.Command{
		Use:   "cliter [TweetContent]",
		Short: "Tweet from cli Command",
		Long:  `You can tweet from cli with cliter command`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%d", len(args))
			msg := strings.Join(args[:], "\n")
			twiReq := &pb.SendTweetRequest{
				AccessToken:  oauth.AuthInfo.AccessToken,
				AccessSecret: oauth.AuthInfo.AccessSecret,
				TweetContent: msg,
			}
			fmt.Printf("Req: %v", twiReq)
			twiRes, err := twitterClient.SendTweet(ctx, twiReq)
			if err != nil {
				log.Fatalf("Failed to Tweet: %v", err)
				return
			}
			if twiRes.IsSuccess {
				log.Printf("Tweet Success")
			}
			return
		},
	}
)

func init() {
	// Dial to server
	baseCtx := context.Background()
	ctx, cancel = context.WithCancel(baseCtx)
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := grpc.DialContext(ctx, "cliter.im-neko.dev:443", grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	if oauth.IsNeedOAuth {
		err = oauth.StartOAuth(ctx, conn)
		if err != nil {
			log.Fatalf("Failed to OAuth: %v", err)
		}
	}
	// Create Client
	twitterClient = pb.NewTweetServiceClient(conn)

}

func main() {
	defer cancel()
	rootCmd.Execute()
	return
}
