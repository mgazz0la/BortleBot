// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/bortles.proto

package bortlespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BortlesService_HelloWorld_FullMethodName = "/bortlespb.BortlesService/HelloWorld"
)

// BortlesServiceClient is the client API for BortlesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BortlesServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type bortlesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBortlesServiceClient(cc grpc.ClientConnInterface) BortlesServiceClient {
	return &bortlesServiceClient{cc}
}

func (c *bortlesServiceClient) HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, BortlesService_HelloWorld_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BortlesServiceServer is the server API for BortlesService service.
// All implementations must embed UnimplementedBortlesServiceServer
// for forward compatibility
type BortlesServiceServer interface {
	HelloWorld(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedBortlesServiceServer()
}

// UnimplementedBortlesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBortlesServiceServer struct {
}

func (UnimplementedBortlesServiceServer) HelloWorld(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedBortlesServiceServer) mustEmbedUnimplementedBortlesServiceServer() {}

// UnsafeBortlesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BortlesServiceServer will
// result in compilation errors.
type UnsafeBortlesServiceServer interface {
	mustEmbedUnimplementedBortlesServiceServer()
}

func RegisterBortlesServiceServer(s grpc.ServiceRegistrar, srv BortlesServiceServer) {
	s.RegisterService(&BortlesService_ServiceDesc, srv)
}

func _BortlesService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BortlesServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BortlesService_HelloWorld_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BortlesServiceServer).HelloWorld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BortlesService_ServiceDesc is the grpc.ServiceDesc for BortlesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BortlesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bortlespb.BortlesService",
	HandlerType: (*BortlesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _BortlesService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bortles.proto",
}
