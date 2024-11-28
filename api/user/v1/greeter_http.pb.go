// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v4.25.1
// source: user/v1/greeter.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGreeterSayHello = "/user.v1.Greeter/SayHello"

type GreeterHTTPServer interface {
	// SayHello Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterHTTPServer(s *http.Server, srv GreeterHTTPServer) {
	r := s.Route("/")
	r.GET("/helloworld/{name}", _Greeter_SayHello0_HTTP_Handler(srv))
}

func _Greeter_SayHello0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type GreeterHTTPClient interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type GreeterHTTPClientImpl struct {
	cc *http.Client
}

func NewGreeterHTTPClient(client *http.Client) GreeterHTTPClient {
	return &GreeterHTTPClientImpl{client}
}

func (c *GreeterHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

const OperationUserSvcCreateDir = "/user.v1.UserSvc/CreateDir"
const OperationUserSvcCreateShare = "/user.v1.UserSvc/CreateShare"
const OperationUserSvcDeleteFiles = "/user.v1.UserSvc/DeleteFiles"
const OperationUserSvcDownload = "/user.v1.UserSvc/Download"
const OperationUserSvcGetShare = "/user.v1.UserSvc/GetShare"
const OperationUserSvcListFiles = "/user.v1.UserSvc/ListFiles"
const OperationUserSvcLogin = "/user.v1.UserSvc/Login"
const OperationUserSvcRegister = "/user.v1.UserSvc/Register"
const OperationUserSvcSaveShare = "/user.v1.UserSvc/SaveShare"
const OperationUserSvcUpload = "/user.v1.UserSvc/Upload"
const OperationUserSvcUploadOver = "/user.v1.UserSvc/UploadOver"

type UserSvcHTTPServer interface {
	CreateDir(context.Context, *CreateDirRequest) (*CommonReply, error)
	CreateShare(context.Context, *CreateShareRequest) (*CommonReply, error)
	DeleteFiles(context.Context, *DeleteFilesRequest) (*CommonReply, error)
	// Downloadrpc Download(DownloadRequest) returns (stream DownloadReply) {
	Download(context.Context, *DownloadRequest) (*DownloadReply, error)
	GetShare(context.Context, *GetShareRequest) (*CommonReply, error)
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesReply, error)
	Login(context.Context, *LoginRequest) (*CommonReply, error)
	Register(context.Context, *RegisterRequest) (*CommonReply, error)
	SaveShare(context.Context, *SaveShareRequest) (*CommonReply, error)
	// Upload rpc GetUser (GetUserRequest) returns (GetUserReply) {
	//   option (google.api.http) = {
	//     get: "/user/{id}"
	//   };
	// }
	Upload(context.Context, *UploadRequest) (*CommonReply, error)
	UploadOver(context.Context, *UploadRequest) (*CommonReply, error)
}

func RegisterUserSvcHTTPServer(s *http.Server, srv UserSvcHTTPServer) {
	r := s.Route("/")
	r.POST("/register", _UserSvc_Register0_HTTP_Handler(srv))
	r.POST("/login", _UserSvc_Login0_HTTP_Handler(srv))
	r.POST("/upload", _UserSvc_Upload0_HTTP_Handler(srv))
	r.POST("/uploadOver", _UserSvc_UploadOver0_HTTP_Handler(srv))
	r.GET("/download/{filename}", _UserSvc_Download0_HTTP_Handler(srv))
	r.GET("/listFiles", _UserSvc_ListFiles0_HTTP_Handler(srv))
	r.POST("/deleteFiles", _UserSvc_DeleteFiles0_HTTP_Handler(srv))
	r.POST("/createDir", _UserSvc_CreateDir0_HTTP_Handler(srv))
	r.POST("/createShare", _UserSvc_CreateShare0_HTTP_Handler(srv))
	r.GET("/getShare/{id}", _UserSvc_GetShare0_HTTP_Handler(srv))
	r.POST("/saveShare", _UserSvc_SaveShare0_HTTP_Handler(srv))
}

func _UserSvc_Register0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_Login0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_Upload0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Upload(ctx, req.(*UploadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_UploadOver0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcUploadOver)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadOver(ctx, req.(*UploadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_Download0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcDownload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Download(ctx, req.(*DownloadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_ListFiles0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFilesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcListFiles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFiles(ctx, req.(*ListFilesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListFilesReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_DeleteFiles0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFilesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcDeleteFiles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFiles(ctx, req.(*DeleteFilesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_CreateDir0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcCreateDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDir(ctx, req.(*CreateDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_CreateShare0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateShareRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcCreateShare)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateShare(ctx, req.(*CreateShareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_GetShare0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetShareRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcGetShare)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetShare(ctx, req.(*GetShareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

func _UserSvc_SaveShare0_HTTP_Handler(srv UserSvcHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveShareRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSvcSaveShare)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveShare(ctx, req.(*SaveShareRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonReply)
		return ctx.Result(200, reply)
	}
}

type UserSvcHTTPClient interface {
	CreateDir(ctx context.Context, req *CreateDirRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	CreateShare(ctx context.Context, req *CreateShareRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	DeleteFiles(ctx context.Context, req *DeleteFilesRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	Download(ctx context.Context, req *DownloadRequest, opts ...http.CallOption) (rsp *DownloadReply, err error)
	GetShare(ctx context.Context, req *GetShareRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	ListFiles(ctx context.Context, req *ListFilesRequest, opts ...http.CallOption) (rsp *ListFilesReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	SaveShare(ctx context.Context, req *SaveShareRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	Upload(ctx context.Context, req *UploadRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
	UploadOver(ctx context.Context, req *UploadRequest, opts ...http.CallOption) (rsp *CommonReply, err error)
}

type UserSvcHTTPClientImpl struct {
	cc *http.Client
}

func NewUserSvcHTTPClient(client *http.Client) UserSvcHTTPClient {
	return &UserSvcHTTPClientImpl{client}
}

func (c *UserSvcHTTPClientImpl) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/createDir"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcCreateDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) CreateShare(ctx context.Context, in *CreateShareRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/createShare"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcCreateShare))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) DeleteFiles(ctx context.Context, in *DeleteFilesRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/deleteFiles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcDeleteFiles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) Download(ctx context.Context, in *DownloadRequest, opts ...http.CallOption) (*DownloadReply, error) {
	var out DownloadReply
	pattern := "/download/{filename}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserSvcDownload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) GetShare(ctx context.Context, in *GetShareRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/getShare/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserSvcGetShare))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...http.CallOption) (*ListFilesReply, error) {
	var out ListFilesReply
	pattern := "/listFiles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserSvcListFiles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) SaveShare(ctx context.Context, in *SaveShareRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/saveShare"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcSaveShare))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) Upload(ctx context.Context, in *UploadRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcUpload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserSvcHTTPClientImpl) UploadOver(ctx context.Context, in *UploadRequest, opts ...http.CallOption) (*CommonReply, error) {
	var out CommonReply
	pattern := "/uploadOver"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSvcUploadOver))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
