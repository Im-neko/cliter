// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tweetpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ReceivePIN(ctx context.Context, in *ReceivePINRequest, opts ...grpc.CallOption) (*ReceivePINResponse, error)
}

type oAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthServiceClient(cc grpc.ClientConnInterface) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/tweet.OAuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) ReceivePIN(ctx context.Context, in *ReceivePINRequest, opts ...grpc.CallOption) (*ReceivePINResponse, error) {
	out := new(ReceivePINResponse)
	err := c.cc.Invoke(ctx, "/tweet.OAuthService/ReceivePIN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
// All implementations must embed UnimplementedOAuthServiceServer
// for forward compatibility
type OAuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ReceivePIN(context.Context, *ReceivePINRequest) (*ReceivePINResponse, error)
	mustEmbedUnimplementedOAuthServiceServer()
}

// UnimplementedOAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (UnimplementedOAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedOAuthServiceServer) ReceivePIN(context.Context, *ReceivePINRequest) (*ReceivePINResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivePIN not implemented")
}
func (UnimplementedOAuthServiceServer) mustEmbedUnimplementedOAuthServiceServer() {}

// UnsafeOAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthServiceServer will
// result in compilation errors.
type UnsafeOAuthServiceServer interface {
	mustEmbedUnimplementedOAuthServiceServer()
}

func RegisterOAuthServiceServer(s grpc.ServiceRegistrar, srv OAuthServiceServer) {
	s.RegisterService(&_OAuthService_serviceDesc, srv)
}

func _OAuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.OAuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_ReceivePIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceivePINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).ReceivePIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.OAuthService/ReceivePIN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).ReceivePIN(ctx, req.(*ReceivePINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tweet.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _OAuthService_Login_Handler,
		},
		{
			MethodName: "ReceivePIN",
			Handler:    _OAuthService_ReceivePIN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tweet.proto",
}

// TweetServiceClient is the client API for TweetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TweetServiceClient interface {
	SendTweet(ctx context.Context, in *SendTweetRequest, opts ...grpc.CallOption) (*SendTweetResponse, error)
}

type tweetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetServiceClient(cc grpc.ClientConnInterface) TweetServiceClient {
	return &tweetServiceClient{cc}
}

func (c *tweetServiceClient) SendTweet(ctx context.Context, in *SendTweetRequest, opts ...grpc.CallOption) (*SendTweetResponse, error) {
	out := new(SendTweetResponse)
	err := c.cc.Invoke(ctx, "/tweet.TweetService/SendTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetServiceServer is the server API for TweetService service.
// All implementations must embed UnimplementedTweetServiceServer
// for forward compatibility
type TweetServiceServer interface {
	SendTweet(context.Context, *SendTweetRequest) (*SendTweetResponse, error)
	mustEmbedUnimplementedTweetServiceServer()
}

// UnimplementedTweetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTweetServiceServer struct {
}

func (UnimplementedTweetServiceServer) SendTweet(context.Context, *SendTweetRequest) (*SendTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTweet not implemented")
}
func (UnimplementedTweetServiceServer) mustEmbedUnimplementedTweetServiceServer() {}

// UnsafeTweetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetServiceServer will
// result in compilation errors.
type UnsafeTweetServiceServer interface {
	mustEmbedUnimplementedTweetServiceServer()
}

func RegisterTweetServiceServer(s grpc.ServiceRegistrar, srv TweetServiceServer) {
	s.RegisterService(&_TweetService_serviceDesc, srv)
}

func _TweetService_SendTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).SendTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tweet.TweetService/SendTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).SendTweet(ctx, req.(*SendTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TweetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tweet.TweetService",
	HandlerType: (*TweetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTweet",
			Handler:    _TweetService_SendTweet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tweet.proto",
}