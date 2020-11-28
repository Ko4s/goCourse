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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	SayManyHello(ctx context.Context, in *GreetManyRequest, opts ...grpc.CallOption) (GreetService_SayManyHelloClient, error)
	GreetManyUsers(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetManyUsersClient, error)
	GreetManyTimes(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	MatchNameWithData(ctx context.Context, in *MatchNameWithDataRequest, opts ...grpc.CallOption) (*MatchNameWithDataResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayManyHello(ctx context.Context, in *GreetManyRequest, opts ...grpc.CallOption) (GreetService_SayManyHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/SayManyHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayManyHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayManyHelloClient interface {
	Recv() (*GreetManyResponse, error)
	grpc.ClientStream
}

type greetServiceSayManyHelloClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayManyHelloClient) Recv() (*GreetManyResponse, error) {
	m := new(GreetManyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetManyUsers(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetManyUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/GreetManyUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyUsersClient{stream}
	return x, nil
}

type GreetService_GreetManyUsersClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyUsersClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyUsersClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetManyUsersClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) MatchNameWithData(ctx context.Context, in *MatchNameWithDataRequest, opts ...grpc.CallOption) (*MatchNameWithDataResponse, error) {
	out := new(MatchNameWithDataResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/MatchNameWithData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *GreetRequest) (*GreetResponse, error)
	SayManyHello(*GreetManyRequest, GreetService_SayManyHelloServer) error
	GreetManyUsers(GreetService_GreetManyUsersServer) error
	GreetManyTimes(GreetService_GreetManyTimesServer) error
	MatchNameWithData(context.Context, *MatchNameWithDataRequest) (*MatchNameWithDataResponse, error)
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayManyHello(*GreetManyRequest, GreetService_SayManyHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayManyHello not implemented")
}
func (UnimplementedGreetServiceServer) GreetManyUsers(GreetService_GreetManyUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyUsers not implemented")
}
func (UnimplementedGreetServiceServer) GreetManyTimes(GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (UnimplementedGreetServiceServer) MatchNameWithData(context.Context, *MatchNameWithDataRequest) (*MatchNameWithDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchNameWithData not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayManyHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayManyHello(m, &greetServiceSayManyHelloServer{stream})
}

type GreetService_SayManyHelloServer interface {
	Send(*GreetManyResponse) error
	grpc.ServerStream
}

type greetServiceSayManyHelloServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayManyHelloServer) Send(m *GreetManyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_GreetManyUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetManyUsers(&greetServiceGreetManyUsersServer{stream})
}

type GreetService_GreetManyUsersServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGreetManyUsersServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyUsersServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetManyUsersServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetManyTimes(&greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetManyTimesServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_MatchNameWithData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchNameWithDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).MatchNameWithData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/MatchNameWithData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).MatchNameWithData(ctx, req.(*MatchNameWithDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
		{
			MethodName: "MatchNameWithData",
			Handler:    _GreetService_MatchNameWithData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayManyHello",
			Handler:       _GreetService_SayManyHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetManyUsers",
			Handler:       _GreetService_GreetManyUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto_files/greet.proto",
}
