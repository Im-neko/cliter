package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/im-neko/cliter/cmd/oauth"
	"github.com/im-neko/cliter/logger"

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
		Long: `You can tweet from cli with cliter command
Example: 
  cliter "Hello Twitter" -- When the tweet contains space, please wrap with double quote
  cliter "line1" line2 "line3" -- To Tweet multiline, please join these line with spacee
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("No tweet text.\nTo tweet \"Hello Twitter\", Enter a command which like: cliter \"Hello Twitter\" ")
				return
			}
			msg := strings.Join(args[:], "\n")
			twiReq := &pb.SendTweetRequest{
				AccessToken:  oauth.AuthInfo.AccessToken,
				AccessSecret: oauth.AuthInfo.AccessSecret,
				TweetContent: msg,
			}
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
    ServerName: "cliter.im-neko.net",
	}
	conn, err := grpc.DialContext(
		ctx,
		"cliter.im-neko.net:443",
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
	)
	if err != nil {
		logger.Error("Could not connect: %v", err)
	}

	if oauth.IsNeedOAuth {
		err = oauth.StartOAuth(ctx, conn)
		if err != nil {
			logger.Error("Failed to OAuth: %v", err)
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
