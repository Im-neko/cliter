# cliter
CLI Twitter Client

# Requirements
Twitter ConsumerKey/Secret

# Getting start
## Build environment
- `go install`
- `export CK=[TwitterConsumerKey]`
- `export CS=[TwitterConsumerSecret]`
## Launch Server
- `cd api && go run main.go`
## Launch Client
In first time, move the authorization page and Tweet will send.
next time, Tweet will send without authorize.
- `cd cmd && go run main.go 'TweetContent'`
