package biz

import (
	"context"
	"driver-back/public"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-playground/validator/v10"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *public.User) (*public.User, error)
	GetUser(ctx context.Context, user *public.User) (*public.User, error)
	CreateFile(ctx context.Context, file *public.File) (*public.File, error)
	CreateFileMinIO(ctx context.Context, path string, file *public.File) error
	CreateFileLocal(ctx context.Context, path_ string, file *public.File) error
	ListFiles(ctx context.Context, uid int64, path string, page, pageSize uint32) (int64, []public.File, error)
	GetFile(ctx context.Context, uid int64, path, filename string) ([]byte, error)
	DeleteFiles(ctx context.Context, uid int64, path string, files []string)
	GetFid(ctx context.Context, uid int64, path_ string, name string) (int64, error)
	CreateShareRecode(ctx context.Context, now time.Time, user *public.User, fids []int64, pwd, exp string) (string, error)
	GetShares(ctx context.Context, id, pwd string) (string, []public.File, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUseCase) CreateUserCase(ctx context.Context, user *public.User) (*public.User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *public.User) (*public.User, error) {
	u.log.Infof("user: %#v", user)
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, err
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) GetUser(ctx context.Context, user *public.User) (*public.User, error) {
	fmt.Printf("user: %#v\n", user)
	if user.Email == "" || user.Password == "" {
		return nil, public.ErrUserNotFound
	}
	return u.repo.GetUser(ctx, user)
}

func (u *UserUseCase) CreateFile(ctx context.Context, path string, file *public.File) (*public.File, error) {
	if file.Type == public.Directory.String() {
		log.Infof("create directory %s for user %d", file.Path, file.UID)
	} else {
		switch public.StoreType {
		case "minio":
			u.repo.CreateFileMinIO(ctx, path, file)
		case "local":
			u.repo.CreateFileLocal(ctx, path, file)
		default:
			return nil, fmt.Errorf("unknown store type")
		}
	}
	return u.repo.CreateFile(ctx, file)
}

func (u *UserUseCase) ListFiles(ctx context.Context, uid int64, path string, page, pageSize uint32) (int64, []public.File, error) {
	return u.repo.ListFiles(ctx, uid, path, page, pageSize)
}

func (u *UserUseCase) GetFile(ctx context.Context, uid int64, path, filename string) ([]byte, error) {
	return u.repo.GetFile(ctx, uid, path, filename)
}

func (u *UserUseCase) DeleteFiles(ctx context.Context, uid int64, path string, files []string) {
	// maybe we no need to delete source file
	// switch public.StoreType {
	// case "minio":
	// 	u.repo.DeleteFilesMinio(ctx, path, file)
	// case "local":
	// 	u.repo.DeleteFilesLocal(ctx, path, file)
	// default:
	// 	return nil, fmt.Errorf("unknown store type")
	// }
	u.repo.DeleteFiles(ctx, uid, path, files)
	//return u.repo.DeleteFiles(ctx, uid, path, files)
}

func (u *UserUseCase) CreateShare(ctx context.Context, user *public.User, names []string, dir, pwd, exp string) (string, error) {
	now := time.Now()
	//TODO: if time.Now() > exp
	var fids []int64
	for _, name := range names {
		if id, err := u.repo.GetFid(ctx, user.ID, dir, name); err == nil {
			fids = append(fids, id)
		}
	}
	return u.repo.CreateShareRecode(ctx, now, user, fids, pwd, exp)
}

func (u *UserUseCase) GetShare(ctx context.Context, sid, pwd string) (string, []public.File, error) {
	return u.repo.GetShares(ctx, sid, pwd)
}
