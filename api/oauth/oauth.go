package oauth

import (
	"github.com/dghubble/oauth1"
)

// AccessInfo is contain access information
type AccessInfo struct {
	AccessToken  string `json:"at,string"`
	AccessSecret string `json:"as,string"`
}

// LoginImpl is login imple
func LoginImpl(config *oauth1.Config) (requestToken string, url string, err error) {
	requestToken, _, err = config.RequestToken()
	if err != nil {
		return "", "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}
	if err != nil {
		// ignore error
	}
	return requestToken, authorizationURL.String(), err
}

// ReceivePinImpl is ReceivePinImpl
func ReceivePinImpl(config *oauth1.Config, requestToken string, verifier string) (*AccessInfo, error) {
	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return &AccessInfo{accessToken, accessSecret}, err
}
