// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service2pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// Service2Client is the client API for Service2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service2Client interface {
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error)
}

type service2Client struct {
	cc grpc.ClientConnInterface
}

func NewService2Client(cc grpc.ClientConnInterface) Service2Client {
	return &service2Client{cc}
}

func (c *service2Client) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error) {
	out := new(HiResponse)
	err := c.cc.Invoke(ctx, "/service2pb.Service2/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service2Server is the server API for Service2 service.
// All implementations must embed UnimplementedService2Server
// for forward compatibility
type Service2Server interface {
	SayHi(context.Context, *HiRequest) (*HiResponse, error)
	mustEmbedUnimplementedService2Server()
}

// UnimplementedService2Server must be embedded to have forward compatible implementations.
type UnimplementedService2Server struct {
}

func (UnimplementedService2Server) SayHi(context.Context, *HiRequest) (*HiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedService2Server) mustEmbedUnimplementedService2Server() {}

// UnsafeService2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Service2Server will
// result in compilation errors.
type UnsafeService2Server interface {
	mustEmbedUnimplementedService2Server()
}

func RegisterService2Server(s *grpc.Server, srv Service2Server) {
	s.RegisterService(&_Service2_serviceDesc, srv)
}

func _Service2_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service2Server).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service2pb.Service2/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service2Server).SayHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service2pb.Service2",
	HandlerType: (*Service2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Service2_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
