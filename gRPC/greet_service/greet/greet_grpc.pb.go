// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetClient interface {
	SayHello(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetManyUsers(ctx context.Context, opts ...grpc.CallOption) (Greet_GreetManyUsersClient, error)
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (Greet_GreetManyTimesClient, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) SayHello(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.Greet/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetClient) GreetManyUsers(ctx context.Context, opts ...grpc.CallOption) (Greet_GreetManyUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greet_serviceDesc.Streams[0], "/greet.Greet/GreetManyUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetGreetManyUsersClient{stream}
	return x, nil
}

type Greet_GreetManyUsersClient interface {
	Send(*GreetManyUsersRequest) error
	CloseAndRecv() (*GreetManyUsersResponse, error)
	grpc.ClientStream
}

type greetGreetManyUsersClient struct {
	grpc.ClientStream
}

func (x *greetGreetManyUsersClient) Send(m *GreetManyUsersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetGreetManyUsersClient) CloseAndRecv() (*GreetManyUsersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetManyUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (Greet_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greet_serviceDesc.Streams[1], "/greet.Greet/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greet_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServer is the server API for Greet service.
// All implementations must embed UnimplementedGreetServer
// for forward compatibility
type GreetServer interface {
	SayHello(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetManyUsers(Greet_GreetManyUsersServer) error
	GreetManyTimes(*GreetManyTimesRequest, Greet_GreetManyTimesServer) error
	mustEmbedUnimplementedGreetServer()
}

// UnimplementedGreetServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) SayHello(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServer) GreetManyUsers(Greet_GreetManyUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyUsers not implemented")
}
func (UnimplementedGreetServer) GreetManyTimes(*GreetManyTimesRequest, Greet_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (UnimplementedGreetServer) mustEmbedUnimplementedGreetServer() {}

// UnsafeGreetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServer will
// result in compilation errors.
type UnsafeGreetServer interface {
	mustEmbedUnimplementedGreetServer()
}

func RegisterGreetServer(s grpc.ServiceRegistrar, srv GreetServer) {
	s.RegisterService(&_Greet_serviceDesc, srv)
}

func _Greet_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.Greet/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).SayHello(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greet_GreetManyUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServer).GreetManyUsers(&greetGreetManyUsersServer{stream})
}

type Greet_GreetManyUsersServer interface {
	SendAndClose(*GreetManyUsersResponse) error
	Recv() (*GreetManyUsersRequest, error)
	grpc.ServerStream
}

type greetGreetManyUsersServer struct {
	grpc.ServerStream
}

func (x *greetGreetManyUsersServer) SendAndClose(m *GreetManyUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetGreetManyUsersServer) Recv() (*GreetManyUsersRequest, error) {
	m := new(GreetManyUsersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greet_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServer).GreetManyTimes(m, &greetGreetManyTimesServer{stream})
}

type Greet_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Greet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greet_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyUsers",
			Handler:       _Greet_GreetManyUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetManyTimes",
			Handler:       _Greet_GreetManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto_files/greet.proto",
}
