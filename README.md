# cliter
CLI Simple Twitter Client
- Tweet Only  
![demo](https://im-neko.dev/assets/cliter-demo.gif)
 
## How to install
`brew tap im-neko/tap`  
`brew install im-neko/tap/cliterk`

## How to use
### Simple Tweet 
`cliter HelloTwitter`

### Multiline Tweet
`cliter Hello Twitter`

### Tweet which includes spaces
`cliter "Hello Twitter"`

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
