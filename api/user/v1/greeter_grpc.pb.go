// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: user/v1/greeter.proto

package v1

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
	Greeter_SayHello_FullMethodName = "/user.v1.Greeter/SayHello"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/greeter.proto",
}

const (
	UserSvc_Register_FullMethodName    = "/user.v1.UserSvc/Register"
	UserSvc_Login_FullMethodName       = "/user.v1.UserSvc/Login"
	UserSvc_Upload_FullMethodName      = "/user.v1.UserSvc/Upload"
	UserSvc_UploadOver_FullMethodName  = "/user.v1.UserSvc/UploadOver"
	UserSvc_Download_FullMethodName    = "/user.v1.UserSvc/Download"
	UserSvc_ListFiles_FullMethodName   = "/user.v1.UserSvc/ListFiles"
	UserSvc_DeleteFiles_FullMethodName = "/user.v1.UserSvc/DeleteFiles"
)

// UserSvcClient is the client API for UserSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSvcClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*CommonReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CommonReply, error)
	//	rpc GetUser (GetUserRequest) returns (GetUserReply) {
	//	  option (google.api.http) = {
	//	    get: "/user/{id}"
	//	  };
	//	}
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*CommonReply, error)
	UploadOver(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*CommonReply, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadReply, error)
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesReply, error)
	DeleteFiles(ctx context.Context, in *DeleteFilesRequest, opts ...grpc.CallOption) (*CommonReply, error)
}

type userSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSvcClient(cc grpc.ClientConnInterface) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, UserSvc_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, UserSvc_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, UserSvc_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UploadOver(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, UserSvc_UploadOver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadReply, error) {
	out := new(DownloadReply)
	err := c.cc.Invoke(ctx, UserSvc_Download_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesReply, error) {
	out := new(ListFilesReply)
	err := c.cc.Invoke(ctx, UserSvc_ListFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) DeleteFiles(ctx context.Context, in *DeleteFilesRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, UserSvc_DeleteFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSvcServer is the server API for UserSvc service.
// All implementations must embed UnimplementedUserSvcServer
// for forward compatibility
type UserSvcServer interface {
	Register(context.Context, *RegisterRequest) (*CommonReply, error)
	Login(context.Context, *LoginRequest) (*CommonReply, error)
	//	rpc GetUser (GetUserRequest) returns (GetUserReply) {
	//	  option (google.api.http) = {
	//	    get: "/user/{id}"
	//	  };
	//	}
	Upload(context.Context, *UploadRequest) (*CommonReply, error)
	UploadOver(context.Context, *UploadRequest) (*CommonReply, error)
	Download(context.Context, *DownloadRequest) (*DownloadReply, error)
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesReply, error)
	DeleteFiles(context.Context, *DeleteFilesRequest) (*CommonReply, error)
	mustEmbedUnimplementedUserSvcServer()
}

// UnimplementedUserSvcServer must be embedded to have forward compatible implementations.
type UnimplementedUserSvcServer struct {
}

func (UnimplementedUserSvcServer) Register(context.Context, *RegisterRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserSvcServer) Login(context.Context, *LoginRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserSvcServer) Upload(context.Context, *UploadRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUserSvcServer) UploadOver(context.Context, *UploadRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadOver not implemented")
}
func (UnimplementedUserSvcServer) Download(context.Context, *DownloadRequest) (*DownloadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedUserSvcServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedUserSvcServer) DeleteFiles(context.Context, *DeleteFilesRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFiles not implemented")
}
func (UnimplementedUserSvcServer) mustEmbedUnimplementedUserSvcServer() {}

// UnsafeUserSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSvcServer will
// result in compilation errors.
type UnsafeUserSvcServer interface {
	mustEmbedUnimplementedUserSvcServer()
}

func RegisterUserSvcServer(s grpc.ServiceRegistrar, srv UserSvcServer) {
	s.RegisterService(&UserSvc_ServiceDesc, srv)
}

func _UserSvc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UploadOver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UploadOver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_UploadOver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UploadOver(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_Download_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_DeleteFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).DeleteFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserSvc_DeleteFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).DeleteFiles(ctx, req.(*DeleteFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSvc_ServiceDesc is the grpc.ServiceDesc for UserSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserSvc_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserSvc_Login_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _UserSvc_Upload_Handler,
		},
		{
			MethodName: "UploadOver",
			Handler:    _UserSvc_UploadOver_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _UserSvc_Download_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _UserSvc_ListFiles_Handler,
		},
		{
			MethodName: "DeleteFiles",
			Handler:    _UserSvc_DeleteFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/greeter.proto",
}