// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: coordinator.proto

package __

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
	Coordinator_AvailablePackages_FullMethodName      = "/pb.Coordinator/AvailablePackages"
	Coordinator_CoordinatorViewPackage_FullMethodName = "/pb.Coordinator/CoordinatorViewPackage"
)

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	AvailablePackages(ctx context.Context, in *View, opts ...grpc.CallOption) (*PackagesResponce, error)
	CoordinatorViewPackage(ctx context.Context, in *View, opts ...grpc.CallOption) (*Package, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) AvailablePackages(ctx context.Context, in *View, opts ...grpc.CallOption) (*PackagesResponce, error) {
	out := new(PackagesResponce)
	err := c.cc.Invoke(ctx, Coordinator_AvailablePackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorViewPackage(ctx context.Context, in *View, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
// All implementations must embed UnimplementedCoordinatorServer
// for forward compatibility
type CoordinatorServer interface {
	AvailablePackages(context.Context, *View) (*PackagesResponce, error)
	CoordinatorViewPackage(context.Context, *View) (*Package, error)
	mustEmbedUnimplementedCoordinatorServer()
}

// UnimplementedCoordinatorServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (UnimplementedCoordinatorServer) AvailablePackages(context.Context, *View) (*PackagesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailablePackages not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorViewPackage(context.Context, *View) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewPackage not implemented")
}
func (UnimplementedCoordinatorServer) mustEmbedUnimplementedCoordinatorServer() {}

// UnsafeCoordinatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServer will
// result in compilation errors.
type UnsafeCoordinatorServer interface {
	mustEmbedUnimplementedCoordinatorServer()
}

func RegisterCoordinatorServer(s grpc.ServiceRegistrar, srv CoordinatorServer) {
	s.RegisterService(&Coordinator_ServiceDesc, srv)
}

func _Coordinator_AvailablePackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).AvailablePackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_AvailablePackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).AvailablePackages(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorViewPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewPackage(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

// Coordinator_ServiceDesc is the grpc.ServiceDesc for Coordinator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coordinator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AvailablePackages",
			Handler:    _Coordinator_AvailablePackages_Handler,
		},
		{
			MethodName: "CoordinatorViewPackage",
			Handler:    _Coordinator_CoordinatorViewPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}
