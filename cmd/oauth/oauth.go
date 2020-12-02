package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	pb "github.com/im-neko/cliter/proto"
	"google.golang.org/grpc"
)

// AccessInfo is contain access information
type AccessInfo struct {
	AccessToken  string `json:"at,string"`
	AccessSecret string `json:"as,string"`
}

var (
	// AccessFileName is AccessFileName
	// TODO: Ecnrypt
	AccessFileName = os.ExpandEnv("$HOME/.cliter.token")
)

var (
	// AuthInfo is instance of AccessInfo
	AuthInfo *AccessInfo
	// IsNeedOAuth is IsNeedOAuth
	IsNeedOAuth bool = true
)

func init() {
	bytes, err := ioutil.ReadFile(AccessFileName)
	if err != nil {
		fmt.Printf("File Read Error: %v\n", err)
		resp := "y"
		fmt.Printf("Start Authorization? (y|N) default: y\n")
		fmt.Printf("Enter: ")
		_, err = fmt.Scanf("%s", &resp)
		if err != nil {
			log.Fatalf("Prompt Phase: %v", err)
			return
		}
		if resp == "y" {
			return
		}
	}
	json.Unmarshal(bytes, &AuthInfo)
	IsNeedOAuth = false
	return
}

// StartOAuth is starting oauth
func StartOAuth(ctx context.Context, conn *grpc.ClientConn) error {
	log.Printf("Start OAuth")
	cl := pb.NewOAuthServiceClient(conn)
	loginReq := &pb.LoginRequest{}
	loginRes, err := cl.Login(ctx, loginReq)
	if err != nil {
		log.Fatalf("Login Phase: %v", err)
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", loginRes.GetAuthUrl())
	err = exec.Command("open", loginRes.GetAuthUrl()).Start()
	if err != nil {
		// ignore error
	}

	fmt.Printf("Paste your PIN here: ")
	var verifier string
	_, err = fmt.Scanf("%s", &verifier)
	if err != nil {
		log.Fatalf("Prompt Phase: %v", err)
	}

	pinReq := &pb.ReceivePINRequest{
		RequestToken: loginRes.GetRequestToken(),
		Verifier:     verifier,
	}
	pinRes, err := cl.ReceivePIN(ctx, pinReq)
	if err != nil {
		log.Fatalf("PIN Phase: %v", err)
	}

	info := &AccessInfo{
		AccessToken:  pinRes.GetAccessToken(),
		AccessSecret: pinRes.GetAccessSecret(),
	}

	err = writeToken(info)
	if err != nil {
		log.Fatalf("Write Phase: %v", err)
	}

	log.Println("Wrote tokens")

	IsNeedOAuth = false
	AuthInfo = info

	return nil
}

func writeToken(info *AccessInfo) error {
	bytes, _ := json.Marshal(info)

	file, err := os.Create(AccessFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
