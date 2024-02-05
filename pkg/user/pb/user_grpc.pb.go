// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: user.proto

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
	User_UserLoginRequest_FullMethodName         = "/pb.User/UserLoginRequest"
	User_UserSignupRequest_FullMethodName        = "/pb.User/UserSignupRequest"
	User_UserSignupVerifyRequest_FullMethodName  = "/pb.User/UserSignupVerifyRequest"
	User_UserForgetPassword_FullMethodName       = "/pb.User/UserForgetPassword"
	User_UserForgetPasswordVerify_FullMethodName = "/pb.User/UserForgetPasswordVerify"
	User_UserNewPassword_FullMethodName          = "/pb.User/UserNewPassword"
	User_UserProfileUpdate_FullMethodName        = "/pb.User/UserProfileUpdate"
	User_UserViewDestination_FullMethodName      = "/pb.User/UserViewDestination"
	User_UserViewActivity_FullMethodName         = "/pb.User/UserViewActivity"
	User_UserViewCatagories_FullMethodName       = "/pb.User/UserViewCatagories"
	User_UserSearchPacakge_FullMethodName        = "/pb.User/UserSearchPacakge"
	User_UserFilterPackage_FullMethodName        = "/pb.User/UserFilterPackage"
	User_UserViewPackage_FullMethodName          = "/pb.User/UserViewPackage"
	User_UserViewFoodMenu_FullMethodName         = "/pb.User/UserViewFoodMenu"
	User_UserTravellerDetails_FullMethodName     = "/pb.User/UserTravellerDetails"
	User_UserOfflineBooking_FullMethodName       = "/pb.User/UserOfflineBooking"
	User_UserViewPackages_FullMethodName         = "/pb.User/UserViewPackages"
	User_UserOnlinePayment_FullMethodName        = "/pb.User/UserOnlinePayment"
	User_UserPaymentConfirmed_FullMethodName     = "/pb.User/UserPaymentConfirmed"
	User_UserViewHistory_FullMethodName          = "/pb.User/UserViewHistory"
	User_UserViewBooking_FullMethodName          = "/pb.User/UserViewBooking"
	User_UserCancelBooking_FullMethodName        = "/pb.User/UserCancelBooking"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserLoginRequest(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserLoginResponce, error)
	UserSignupRequest(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*UserResponce, error)
	UserSignupVerifyRequest(ctx context.Context, in *UserVerify, opts ...grpc.CallOption) (*UserResponce, error)
	UserForgetPassword(ctx context.Context, in *UserforgetPassword, opts ...grpc.CallOption) (*UserResponce, error)
	UserForgetPasswordVerify(ctx context.Context, in *UserforgetPasswordVerify, opts ...grpc.CallOption) (*UserResponce, error)
	UserNewPassword(ctx context.Context, in *Usernewpassword, opts ...grpc.CallOption) (*UserResponce, error)
	UserProfileUpdate(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*UserResponce, error)
	UserViewDestination(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserDestination, error)
	UserViewActivity(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserActivity, error)
	UserViewCatagories(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserCategories, error)
	UserSearchPacakge(ctx context.Context, in *UserSearch, opts ...grpc.CallOption) (*UserPackages, error)
	UserFilterPackage(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserPackages, error)
	UserViewPackage(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserPackage, error)
	UserViewFoodMenu(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserFoodMenus, error)
	UserTravellerDetails(ctx context.Context, in *UserTravellerRequest, opts ...grpc.CallOption) (*UserTravellerResponse, error)
	UserOfflineBooking(ctx context.Context, in *UserBooking, opts ...grpc.CallOption) (*UserBookingResponce, error)
	UserViewPackages(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserPackages, error)
	UserOnlinePayment(ctx context.Context, in *UserBooking, opts ...grpc.CallOption) (*UserOnlinePaymentResponse, error)
	UserPaymentConfirmed(ctx context.Context, in *UserPaymentConfirmedRequest, opts ...grpc.CallOption) (*UserBookingResponce, error)
	UserViewHistory(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserHistories, error)
	UserViewBooking(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserHistory, error)
	UserCancelBooking(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserResponce, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLoginRequest(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserLoginResponce, error) {
	out := new(UserLoginResponce)
	err := c.cc.Invoke(ctx, User_UserLoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserSignupRequest(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserSignupRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserSignupVerifyRequest(ctx context.Context, in *UserVerify, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserSignupVerifyRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserForgetPassword(ctx context.Context, in *UserforgetPassword, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserForgetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserForgetPasswordVerify(ctx context.Context, in *UserforgetPasswordVerify, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserForgetPasswordVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserNewPassword(ctx context.Context, in *Usernewpassword, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserNewPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserProfileUpdate(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserProfileUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewDestination(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserDestination, error) {
	out := new(UserDestination)
	err := c.cc.Invoke(ctx, User_UserViewDestination_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewActivity(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserActivity, error) {
	out := new(UserActivity)
	err := c.cc.Invoke(ctx, User_UserViewActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewCatagories(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserCategories, error) {
	out := new(UserCategories)
	err := c.cc.Invoke(ctx, User_UserViewCatagories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserSearchPacakge(ctx context.Context, in *UserSearch, opts ...grpc.CallOption) (*UserPackages, error) {
	out := new(UserPackages)
	err := c.cc.Invoke(ctx, User_UserSearchPacakge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserFilterPackage(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*UserPackages, error) {
	out := new(UserPackages)
	err := c.cc.Invoke(ctx, User_UserFilterPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewPackage(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserPackage, error) {
	out := new(UserPackage)
	err := c.cc.Invoke(ctx, User_UserViewPackage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewFoodMenu(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserFoodMenus, error) {
	out := new(UserFoodMenus)
	err := c.cc.Invoke(ctx, User_UserViewFoodMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserTravellerDetails(ctx context.Context, in *UserTravellerRequest, opts ...grpc.CallOption) (*UserTravellerResponse, error) {
	out := new(UserTravellerResponse)
	err := c.cc.Invoke(ctx, User_UserTravellerDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOfflineBooking(ctx context.Context, in *UserBooking, opts ...grpc.CallOption) (*UserBookingResponce, error) {
	out := new(UserBookingResponce)
	err := c.cc.Invoke(ctx, User_UserOfflineBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewPackages(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserPackages, error) {
	out := new(UserPackages)
	err := c.cc.Invoke(ctx, User_UserViewPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOnlinePayment(ctx context.Context, in *UserBooking, opts ...grpc.CallOption) (*UserOnlinePaymentResponse, error) {
	out := new(UserOnlinePaymentResponse)
	err := c.cc.Invoke(ctx, User_UserOnlinePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserPaymentConfirmed(ctx context.Context, in *UserPaymentConfirmedRequest, opts ...grpc.CallOption) (*UserBookingResponce, error) {
	out := new(UserBookingResponce)
	err := c.cc.Invoke(ctx, User_UserPaymentConfirmed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewHistory(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserHistories, error) {
	out := new(UserHistories)
	err := c.cc.Invoke(ctx, User_UserViewHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserViewBooking(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserHistory, error) {
	out := new(UserHistory)
	err := c.cc.Invoke(ctx, User_UserViewBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserCancelBooking(ctx context.Context, in *UserView, opts ...grpc.CallOption) (*UserResponce, error) {
	out := new(UserResponce)
	err := c.cc.Invoke(ctx, User_UserCancelBooking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserLoginRequest(context.Context, *UserLogin) (*UserLoginResponce, error)
	UserSignupRequest(context.Context, *UserSignup) (*UserResponce, error)
	UserSignupVerifyRequest(context.Context, *UserVerify) (*UserResponce, error)
	UserForgetPassword(context.Context, *UserforgetPassword) (*UserResponce, error)
	UserForgetPasswordVerify(context.Context, *UserforgetPasswordVerify) (*UserResponce, error)
	UserNewPassword(context.Context, *Usernewpassword) (*UserResponce, error)
	UserProfileUpdate(context.Context, *UserSignup) (*UserResponce, error)
	UserViewDestination(context.Context, *UserView) (*UserDestination, error)
	UserViewActivity(context.Context, *UserView) (*UserActivity, error)
	UserViewCatagories(context.Context, *UserView) (*UserCategories, error)
	UserSearchPacakge(context.Context, *UserSearch) (*UserPackages, error)
	UserFilterPackage(context.Context, *UserFilter) (*UserPackages, error)
	UserViewPackage(context.Context, *UserView) (*UserPackage, error)
	UserViewFoodMenu(context.Context, *UserView) (*UserFoodMenus, error)
	UserTravellerDetails(context.Context, *UserTravellerRequest) (*UserTravellerResponse, error)
	UserOfflineBooking(context.Context, *UserBooking) (*UserBookingResponce, error)
	UserViewPackages(context.Context, *UserView) (*UserPackages, error)
	UserOnlinePayment(context.Context, *UserBooking) (*UserOnlinePaymentResponse, error)
	UserPaymentConfirmed(context.Context, *UserPaymentConfirmedRequest) (*UserBookingResponce, error)
	UserViewHistory(context.Context, *UserView) (*UserHistories, error)
	UserViewBooking(context.Context, *UserView) (*UserHistory, error)
	UserCancelBooking(context.Context, *UserView) (*UserResponce, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserLoginRequest(context.Context, *UserLogin) (*UserLoginResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginRequest not implemented")
}
func (UnimplementedUserServer) UserSignupRequest(context.Context, *UserSignup) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignupRequest not implemented")
}
func (UnimplementedUserServer) UserSignupVerifyRequest(context.Context, *UserVerify) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignupVerifyRequest not implemented")
}
func (UnimplementedUserServer) UserForgetPassword(context.Context, *UserforgetPassword) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserForgetPassword not implemented")
}
func (UnimplementedUserServer) UserForgetPasswordVerify(context.Context, *UserforgetPasswordVerify) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserForgetPasswordVerify not implemented")
}
func (UnimplementedUserServer) UserNewPassword(context.Context, *Usernewpassword) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserNewPassword not implemented")
}
func (UnimplementedUserServer) UserProfileUpdate(context.Context, *UserSignup) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserProfileUpdate not implemented")
}
func (UnimplementedUserServer) UserViewDestination(context.Context, *UserView) (*UserDestination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewDestination not implemented")
}
func (UnimplementedUserServer) UserViewActivity(context.Context, *UserView) (*UserActivity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewActivity not implemented")
}
func (UnimplementedUserServer) UserViewCatagories(context.Context, *UserView) (*UserCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewCatagories not implemented")
}
func (UnimplementedUserServer) UserSearchPacakge(context.Context, *UserSearch) (*UserPackages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSearchPacakge not implemented")
}
func (UnimplementedUserServer) UserFilterPackage(context.Context, *UserFilter) (*UserPackages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFilterPackage not implemented")
}
func (UnimplementedUserServer) UserViewPackage(context.Context, *UserView) (*UserPackage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewPackage not implemented")
}
func (UnimplementedUserServer) UserViewFoodMenu(context.Context, *UserView) (*UserFoodMenus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewFoodMenu not implemented")
}
func (UnimplementedUserServer) UserTravellerDetails(context.Context, *UserTravellerRequest) (*UserTravellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserTravellerDetails not implemented")
}
func (UnimplementedUserServer) UserOfflineBooking(context.Context, *UserBooking) (*UserBookingResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOfflineBooking not implemented")
}
func (UnimplementedUserServer) UserViewPackages(context.Context, *UserView) (*UserPackages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewPackages not implemented")
}
func (UnimplementedUserServer) UserOnlinePayment(context.Context, *UserBooking) (*UserOnlinePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOnlinePayment not implemented")
}
func (UnimplementedUserServer) UserPaymentConfirmed(context.Context, *UserPaymentConfirmedRequest) (*UserBookingResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserPaymentConfirmed not implemented")
}
func (UnimplementedUserServer) UserViewHistory(context.Context, *UserView) (*UserHistories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewHistory not implemented")
}
func (UnimplementedUserServer) UserViewBooking(context.Context, *UserView) (*UserHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewBooking not implemented")
}
func (UnimplementedUserServer) UserCancelBooking(context.Context, *UserView) (*UserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCancelBooking not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserLoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLoginRequest(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserSignupRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignupRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSignupRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignupRequest(ctx, req.(*UserSignup))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserSignupVerifyRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSignupVerifyRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSignupVerifyRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSignupVerifyRequest(ctx, req.(*UserVerify))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserforgetPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserForgetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserForgetPassword(ctx, req.(*UserforgetPassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserForgetPasswordVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserforgetPasswordVerify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserForgetPasswordVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserForgetPasswordVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserForgetPasswordVerify(ctx, req.(*UserforgetPasswordVerify))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserNewPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Usernewpassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserNewPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserNewPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserNewPassword(ctx, req.(*Usernewpassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserProfileUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserProfileUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserProfileUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserProfileUpdate(ctx, req.(*UserSignup))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewDestination(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewActivity(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewCatagories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewCatagories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewCatagories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewCatagories(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserSearchPacakge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserSearchPacakge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserSearchPacakge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserSearchPacakge(ctx, req.(*UserSearch))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserFilterPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserFilterPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserFilterPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserFilterPackage(ctx, req.(*UserFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewPackage(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewFoodMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewFoodMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewFoodMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewFoodMenu(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserTravellerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTravellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserTravellerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserTravellerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserTravellerDetails(ctx, req.(*UserTravellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserOfflineBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOfflineBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOfflineBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOfflineBooking(ctx, req.(*UserBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewPackages(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserOnlinePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOnlinePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOnlinePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOnlinePayment(ctx, req.(*UserBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserPaymentConfirmed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPaymentConfirmedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserPaymentConfirmed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserPaymentConfirmed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserPaymentConfirmed(ctx, req.(*UserPaymentConfirmedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewHistory(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserViewBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserViewBooking(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserCancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserView)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserCancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserCancelBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserCancelBooking(ctx, req.(*UserView))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLoginRequest",
			Handler:    _User_UserLoginRequest_Handler,
		},
		{
			MethodName: "UserSignupRequest",
			Handler:    _User_UserSignupRequest_Handler,
		},
		{
			MethodName: "UserSignupVerifyRequest",
			Handler:    _User_UserSignupVerifyRequest_Handler,
		},
		{
			MethodName: "UserForgetPassword",
			Handler:    _User_UserForgetPassword_Handler,
		},
		{
			MethodName: "UserForgetPasswordVerify",
			Handler:    _User_UserForgetPasswordVerify_Handler,
		},
		{
			MethodName: "UserNewPassword",
			Handler:    _User_UserNewPassword_Handler,
		},
		{
			MethodName: "UserProfileUpdate",
			Handler:    _User_UserProfileUpdate_Handler,
		},
		{
			MethodName: "UserViewDestination",
			Handler:    _User_UserViewDestination_Handler,
		},
		{
			MethodName: "UserViewActivity",
			Handler:    _User_UserViewActivity_Handler,
		},
		{
			MethodName: "UserViewCatagories",
			Handler:    _User_UserViewCatagories_Handler,
		},
		{
			MethodName: "UserSearchPacakge",
			Handler:    _User_UserSearchPacakge_Handler,
		},
		{
			MethodName: "UserFilterPackage",
			Handler:    _User_UserFilterPackage_Handler,
		},
		{
			MethodName: "UserViewPackage",
			Handler:    _User_UserViewPackage_Handler,
		},
		{
			MethodName: "UserViewFoodMenu",
			Handler:    _User_UserViewFoodMenu_Handler,
		},
		{
			MethodName: "UserTravellerDetails",
			Handler:    _User_UserTravellerDetails_Handler,
		},
		{
			MethodName: "UserOfflineBooking",
			Handler:    _User_UserOfflineBooking_Handler,
		},
		{
			MethodName: "UserViewPackages",
			Handler:    _User_UserViewPackages_Handler,
		},
		{
			MethodName: "UserOnlinePayment",
			Handler:    _User_UserOnlinePayment_Handler,
		},
		{
			MethodName: "UserPaymentConfirmed",
			Handler:    _User_UserPaymentConfirmed_Handler,
		},
		{
			MethodName: "UserViewHistory",
			Handler:    _User_UserViewHistory_Handler,
		},
		{
			MethodName: "UserViewBooking",
			Handler:    _User_UserViewBooking_Handler,
		},
		{
			MethodName: "UserCancelBooking",
			Handler:    _User_UserCancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
