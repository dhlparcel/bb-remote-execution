// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pkg/proto/remoteworker/remoteworker.proto

package remoteworker

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
	OperationQueue_Synchronize_FullMethodName = "/buildbarn.remoteworker.OperationQueue/Synchronize"
)

// OperationQueueClient is the client API for OperationQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationQueueClient interface {
	Synchronize(ctx context.Context, in *SynchronizeRequest, opts ...grpc.CallOption) (*SynchronizeResponse, error)
}

type operationQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationQueueClient(cc grpc.ClientConnInterface) OperationQueueClient {
	return &operationQueueClient{cc}
}

func (c *operationQueueClient) Synchronize(ctx context.Context, in *SynchronizeRequest, opts ...grpc.CallOption) (*SynchronizeResponse, error) {
	out := new(SynchronizeResponse)
	err := c.cc.Invoke(ctx, OperationQueue_Synchronize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationQueueServer is the server API for OperationQueue service.
// All implementations should embed UnimplementedOperationQueueServer
// for forward compatibility
type OperationQueueServer interface {
	Synchronize(context.Context, *SynchronizeRequest) (*SynchronizeResponse, error)
}

// UnimplementedOperationQueueServer should be embedded to have forward compatible implementations.
type UnimplementedOperationQueueServer struct {
}

func (UnimplementedOperationQueueServer) Synchronize(context.Context, *SynchronizeRequest) (*SynchronizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Synchronize not implemented")
}

// UnsafeOperationQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationQueueServer will
// result in compilation errors.
type UnsafeOperationQueueServer interface {
	mustEmbedUnimplementedOperationQueueServer()
}

func RegisterOperationQueueServer(s grpc.ServiceRegistrar, srv OperationQueueServer) {
	s.RegisterService(&OperationQueue_ServiceDesc, srv)
}

func _OperationQueue_Synchronize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynchronizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationQueueServer).Synchronize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperationQueue_Synchronize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationQueueServer).Synchronize(ctx, req.(*SynchronizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationQueue_ServiceDesc is the grpc.ServiceDesc for OperationQueue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationQueue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.remoteworker.OperationQueue",
	HandlerType: (*OperationQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Synchronize",
			Handler:    _OperationQueue_Synchronize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/remoteworker/remoteworker.proto",
}
