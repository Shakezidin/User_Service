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
	Coordinator_AvailablePackages_FullMethodName          = "/pb.Coordinator/AvailablePackages"
	Coordinator_CoordinatorViewPackage_FullMethodName     = "/pb.Coordinator/CoordinatorViewPackage"
	Coordinator_CoordinatorViewDestination_FullMethodName = "/pb.Coordinator/CoordinatorViewDestination"
	Coordinator_CoordinatorViewActivity_FullMethodName    = "/pb.Coordinator/CoordinatorViewActivity"
	Coordinator_ViewCatagories_FullMethodName             = "/pb.Coordinator/ViewCatagories"
	Coordinator_PackageSearch_FullMethodName              = "/pb.Coordinator/PackageSearch"
	Coordinator_TravellerDetails_FullMethodName           = "/pb.Coordinator/TravellerDetails"
)

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorClient interface {
	AvailablePackages(ctx context.Context, in *View, opts ...grpc.CallOption) (*PackagesResponce, error)
	CoordinatorViewPackage(ctx context.Context, in *View, opts ...grpc.CallOption) (*Package, error)
	CoordinatorViewDestination(ctx context.Context, in *View, opts ...grpc.CallOption) (*Destination, error)
	CoordinatorViewActivity(ctx context.Context, in *View, opts ...grpc.CallOption) (*Activity, error)
	ViewCatagories(ctx context.Context, in *View, opts ...grpc.CallOption) (*Catagories, error)
	PackageSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*PackagesResponce, error)
	TravellerDetails(ctx context.Context, in *TravellerRequest, opts ...grpc.CallOption) (*TravellerResponse, error)
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

func (c *coordinatorClient) CoordinatorViewDestination(ctx context.Context, in *View, opts ...grpc.CallOption) (*Destination, error) {
	out := new(Destination)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) CoordinatorViewActivity(ctx context.Context, in *View, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, Coordinator_CoordinatorViewActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) ViewCatagories(ctx context.Context, in *View, opts ...grpc.CallOption) (*Catagories, error) {
	out := new(Catagories)
	err := c.cc.Invoke(ctx, Coordinator_ViewCatagories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) PackageSearch(ctx context.Context, in *Search, opts ...grpc.CallOption) (*PackagesResponce, error) {
	out := new(PackagesResponce)
	err := c.cc.Invoke(ctx, Coordinator_PackageSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) TravellerDetails(ctx context.Context, in *TravellerRequest, opts ...grpc.CallOption) (*TravellerResponse, error) {
	out := new(TravellerResponse)
	err := c.cc.Invoke(ctx, Coordinator_TravellerDetails_FullMethodName, in, out, opts...)
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
	CoordinatorViewDestination(context.Context, *View) (*Destination, error)
	CoordinatorViewActivity(context.Context, *View) (*Activity, error)
	ViewCatagories(context.Context, *View) (*Catagories, error)
	PackageSearch(context.Context, *Search) (*PackagesResponce, error)
	TravellerDetails(context.Context, *TravellerRequest) (*TravellerResponse, error)
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
func (UnimplementedCoordinatorServer) CoordinatorViewDestination(context.Context, *View) (*Destination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewDestination not implemented")
}
func (UnimplementedCoordinatorServer) CoordinatorViewActivity(context.Context, *View) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorViewActivity not implemented")
}
func (UnimplementedCoordinatorServer) ViewCatagories(context.Context, *View) (*Catagories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewCatagories not implemented")
}
func (UnimplementedCoordinatorServer) PackageSearch(context.Context, *Search) (*PackagesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PackageSearch not implemented")
}
func (UnimplementedCoordinatorServer) TravellerDetails(context.Context, *TravellerRequest) (*TravellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TravellerDetails not implemented")
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

func _Coordinator_CoordinatorViewDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewDestination(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_CoordinatorViewActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).CoordinatorViewActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_CoordinatorViewActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).CoordinatorViewActivity(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_ViewCatagories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).ViewCatagories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_ViewCatagories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).ViewCatagories(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_PackageSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).PackageSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_PackageSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).PackageSearch(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_TravellerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TravellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).TravellerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coordinator_TravellerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).TravellerDetails(ctx, req.(*TravellerRequest))
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
		{
			MethodName: "CoordinatorViewDestination",
			Handler:    _Coordinator_CoordinatorViewDestination_Handler,
		},
		{
			MethodName: "CoordinatorViewActivity",
			Handler:    _Coordinator_CoordinatorViewActivity_Handler,
		},
		{
			MethodName: "ViewCatagories",
			Handler:    _Coordinator_ViewCatagories_Handler,
		},
		{
			MethodName: "PackageSearch",
			Handler:    _Coordinator_PackageSearch_Handler,
		},
		{
			MethodName: "TravellerDetails",
			Handler:    _Coordinator_TravellerDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}
