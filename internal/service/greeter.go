package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	v1 "driver-back/api/user/v1"
	"driver-back/internal/biz"
	"driver-back/public"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

type UserService struct {
	v1.UnimplementedUserSvcServer
	uu *biz.UserUseCase
}

func NewUserService(uu *biz.UserUseCase) *UserService {
	return &UserService{uu: uu}
}
func (u *UserService) Register(ctx context.Context, reg *v1.RegisterRequest) (*v1.CommonReply, error) {
	fmt.Printf("++++++++: %v\n", reg)
	user := &public.User{
		Email:    reg.Email,
		Password: reg.Password,
		Phone:    reg.Phone,
		Name:     reg.Username,
	}
	_, err := u.uu.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.CommonReply{Code: 200, Message: user.Name}, nil
}

func getUserFromCtx(ctx context.Context) *public.User {
	tr, _ := transport.FromServerContext(ctx)
	ts := tr.RequestHeader().Get("Authorization")
	if strings.HasPrefix(ts, "Bearer ") {
		ts, _ = strings.CutPrefix(ts, "Bearer ")
	}
	fmt.Println(ts)
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Printf("%#v\n", claim)
		result := make(map[string]interface{}, 2)
		result["user"] = claim["user"]
		fmt.Printf("%#v\n", result)
	}
	user := &public.User{}
	us := claim["user"].(string)
	json.Unmarshal([]byte(us), user)
	return user
}

func (u *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.CommonReply, error) {
	user := &public.User{
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := u.uu.GetUser(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Password = "***************"
	userByte, _ := json.Marshal(user)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = string(userByte)
	// TODO: check if expire func work
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["iat"] = time.Now().Unix()
	tokenStr, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	//getUserFromToken(tokenStr)
	msgMap := map[string]string{
		"token": tokenStr,
		"user":  string(userByte),
	}
	msg, _ := json.Marshal(msgMap)

	return &v1.CommonReply{Code: 200, Message: string(msg)}, nil
}
func (u *UserService) Upload(ctx context.Context, req *v1.UploadRequest) (*v1.CommonReply, error) {
	fn := fmt.Sprintf("%s.part.%d", req.Filename, req.ChunkIndex)
	if src, err := base64.StdEncoding.DecodeString(req.Chunk); err != nil {
		return nil, err
	} else {
		if err := os.WriteFile(path.Join(public.TmpDir, fn), src, os.ModePerm); err != nil {
			return nil, err
		}
	}
	return &v1.CommonReply{Code: 200, Message: fn}, nil
}

func (u *UserService) UploadOver(ctx context.Context, req *v1.UploadRequest) (*v1.CommonReply, error) {
	dir, fn := path.Split(req.Filename)
	fd, err := os.OpenFile(path.Join(public.TmpDir, fn), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}
	size := 0
	h := sha256.New()
	for i := uint32(0); i < req.ChunkIndex; i++ {
		fn := fmt.Sprintf("%s.part.%d", fn, i)
		buf, err := os.ReadFile(path.Join(public.TmpDir, fn))
		if err != nil {
			return nil, err
		}
		n, err := fd.Write(buf)
		if err != nil {
			return nil, err
		}
		size += n
		h.Write(buf)
		os.Remove(path.Join(public.TmpDir, fn))
	}
	hashStr := hex.EncodeToString(h.Sum(nil))
	user := getUserFromCtx(ctx)
	// TODO: identify file type
	u.uu.CreateFile(ctx, fd.Name(), &public.File{
		Hash: hashStr,
		Type: "image",
		UID:  user.ID,
		Name: fn,
		Path: dir,
		Size: int64(size),
		Time: time.Now().Unix(),
	})
	return &v1.CommonReply{Code: 200, Message: "seccess"}, nil
}

func (u *UserService) Download(ctx context.Context, req *v1.DownloadRequest) (*v1.DownloadReply, error) {
	user := getUserFromCtx(ctx)
	content, err := u.uu.GetFile(ctx, user.ID, req.Path, req.Filename)
	if err != nil {
		return nil, nil
	}
	// fmt.Println(string(content))
	return &v1.DownloadReply{Content: content}, nil
	// return &v1.DownloadReply{}, nil
}

func (u *UserService) ListFiles(ctx context.Context, req *v1.ListFilesRequest) (*v1.ListFilesReply, error) {
	user := getUserFromCtx(ctx)
	total, files, err := u.uu.ListFiles(ctx, user.ID, req.Path, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListFilesReply{
		Total:     uint64(total),
		FileInfos: make([]*v1.FileInfo, len(files)),
	}
	for i, v := range files {
		reply.FileInfos[i] = &v1.FileInfo{
			Name: v.Name,
			Size: uint64(v.Size),
			Time: uint64(v.Time),
			Type: v.Type,
		}
	}
	return reply, nil
}

func (u *UserService) DeleteFiles(ctx context.Context, req *v1.DeleteFilesRequest) (*v1.CommonReply, error) {
	user := getUserFromCtx(ctx)
	u.uu.DeleteFiles(ctx, user.ID, req.Path, req.Files)
	return nil, nil
}
