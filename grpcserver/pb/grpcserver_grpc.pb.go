// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ZkServiceClient is the client API for ZkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZkServiceClient interface {
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Message, error)
}

type zkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZkServiceClient(cc grpc.ClientConnInterface) ZkServiceClient {
	return &zkServiceClient{cc}
}

func (c *zkServiceClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/pb.ZkService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZkServiceServer is the server API for ZkService service.
// All implementations must embed UnimplementedZkServiceServer
// for forward compatibility
type ZkServiceServer interface {
	Get(context.Context, *emptypb.Empty) (*Message, error)
	mustEmbedUnimplementedZkServiceServer()
}

// UnimplementedZkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedZkServiceServer struct {
}

func (UnimplementedZkServiceServer) Get(context.Context, *emptypb.Empty) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedZkServiceServer) mustEmbedUnimplementedZkServiceServer() {}

// UnsafeZkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZkServiceServer will
// result in compilation errors.
type UnsafeZkServiceServer interface {
	mustEmbedUnimplementedZkServiceServer()
}

func RegisterZkServiceServer(s grpc.ServiceRegistrar, srv ZkServiceServer) {
	s.RegisterService(&ZkService_ServiceDesc, srv)
}

func _ZkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ZkService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZkServiceServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ZkService_ServiceDesc is the grpc.ServiceDesc for ZkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ZkService",
	HandlerType: (*ZkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ZkService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/grpcserver.proto",
}
