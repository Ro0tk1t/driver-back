package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"driver-back/internal/biz"
	"driver-back/public"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sony/sonyflake"
	"xorm.io/xorm"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}

type userRepo struct {
	data   *Data
	engine *xorm.Engine
	minio  *minio.Client
	log    *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	engine, err := xorm.NewEngine(data.Driver, data.DBAddr)
	if err != nil {
		panic(err)
	}
	engine.Sync(new(public.User))
	if err := engine.Sync(new(public.File)); err != nil {
		fmt.Print(err)
	}
	if err := engine.Sync(new(public.ShareInfo)); err != nil {
		fmt.Print(err)
	}

	client, err := minio.New(data.MinioAccess.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV2(data.MinioAccess.KeyID, data.MinioAccess.KeyPass, ""),
	})
	if err != nil {
		panic(err)
	}
	return &userRepo{
		data:   data,
		engine: engine,
		minio:  client,
		log:    log.NewHelper(logger),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *public.User) (*public.User, error) {
	_, err := u.engine.Insert(user)
	u.log.Debugf("err:::::::::: %#v", err)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) GetUser(ctx context.Context, user *public.User) (*public.User, error) {
	// TODO: pwd md5
	has, err := u.engine.Where("email=?", user.Email).And("password=?", user.Password).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, public.ErrUserNotFound
	}
	return user, nil
}
func (u *userRepo) FindUserByEmail(ctx context.Context, email string) (*public.User, error) {
	user := &public.User{Email: email}
	has, err := u.engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, public.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepo) CreateFile(ctx context.Context, file *public.File) (*public.File, error) {
	has, err := u.engine.Where("name=?", file.Name).And("uid = ?", file.UID).And("path=?", file.Path).Get(&public.File{})
	if err != nil {
		return nil, err
	}
	// FIXME: cover exist file/dir ?
	if has {
		u.log.Error("file already exists ", "hash ", file.Hash)
		return file, nil
	}
	if _, err := u.engine.Insert(file); err != nil {
		return nil, err
	}
	return file, nil
}

func (u *userRepo) CreateFileMinIO(ctx context.Context, path string, file *public.File) error {
	bucket := "z" + file.Hash[:2]
	exists, err := u.minio.BucketExists(context.Background(), bucket)
	if err != nil {
		u.log.Error(err)
		return err
	}
	if !exists {
		err = u.minio.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			u.log.Error(err)
			return err
		}
	}
	_, err = u.minio.StatObject(ctx, bucket, file.Hash, minio.GetObjectOptions{})
	if err != nil {
		info, err := u.minio.FPutObject(context.Background(), bucket, file.Hash[2:], path, minio.PutObjectOptions{})
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
			return err
		}
		fmt.Printf("%#v", info)
	}
	// fmt.Printf("%#v\n", o)
	return nil
}
func (u *userRepo) CreateFileLocal(ctx context.Context, path_ string, file *public.File) error {
	index := file.Hash[:2]
	index = path.Join(public.TmpDir, index)
	if _, err := os.Stat(index); os.IsNotExist(err) {
		if err := os.Mkdir(index, os.ModePerm); err != nil {
			log.Error("index exists: ", index)
		}
	}
	if err := os.Rename(path_, path.Join(index, file.Hash[2:])); err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetFile(ctx context.Context, uid int64, path_, filename string) ([]byte, error) {
	var hash string
	has, err := u.engine.Table(&public.File{}).Cols("hash").Where("uid=?", uid).And("path=?", path_).And("name=?", filename).Get(&hash)
	if err != nil || !has {
		return nil, err
	}
	p := fmt.Sprintf("%s/%s/%s", public.TmpDir, hash[:2], hash[2:])
	if _, err := os.Stat(p); !os.IsNotExist(err) {
		b, _ := os.ReadFile(p)
		return b, nil
	}
	bucket := "z" + hash[:2]
	exists, err := u.minio.BucketExists(context.Background(), bucket)
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	if !exists {
		err = u.minio.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			u.log.Error(err)
			return nil, err
		}
	}
	_, err = u.minio.StatObject(ctx, bucket, hash, minio.GetObjectOptions{})
	if err != nil {
		err := u.minio.FGetObject(context.Background(), bucket, hash[2:], p, minio.GetObjectOptions{})
		// info, err := u.minio.FPutObject(context.Background(), bucket, hash[2:], path, minio.PutObjectOptions{})
		// fmt.Printf("%#v", info)
		if err != nil {
			log.Fatalf("Error get file: %v", err)
			return nil, err
		}
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (u *userRepo) ListFiles(ctx context.Context, uid int64, path string, page, pageSize uint32) (int64, []public.File, error) {
	finfo := &public.File{
		UID:  uid,
		Path: path,
	}
	total, _ := u.engine.Where("uid=?", uid).And("path=?", path).Count(finfo)
	var files []public.File
	if err := u.engine.Where("uid=?", uid).And("path=?", path).Find(&files); err != nil {
		return 0, nil, err
	}
	return total, files, nil
}

func (u *userRepo) DeleteFiles(ctx context.Context, uid int64, path_ string, files []string) {
	fmt.Println(uid, path_, files)
	for _, f := range files {
		fs := public.File{UID: uid, Path: path_, Name: f}
		has, err := u.engine.Get(&fs)
		if !has {
			u.log.Errorf("path: %s, name: %s  not found!", path_, f)
			u.log.Errorf("%#v", err)
			return
		}
		if err != nil {
			u.log.Error("db error, ", err)
			return
		}
		if fs.Type == public.Directory.String() {
			// TODO: use TRANSACTION
			p := path.Join(path_, f) + "/"
			var fss []public.File
			if err := u.engine.Table(&fs).Where("uid=?", uid).And("path=?", p).Cols("name").Find(&fss); err != nil {
				u.log.Error("search file error, ", err)
			}
			for _, fs_ := range fss {
				u.DeleteFiles(ctx, uid, p, []string{fs_.Name})
			}
		}
		if _, err := u.engine.Where("uid=?", uid).And("path=?", path_).And("name=?", f).Delete(&public.File{}); err != nil {
			u.log.Error(err)
		}
	}
}

func (u *userRepo) GetFid(ctx context.Context, uid int64, path_ string, name string) (int64, error) {
	file := public.File{}
	has, err := u.engine.Where("uid=?", uid).And("path=?", path_).And("name=?", name).Get(&file)
	if !has {
		return 0, fmt.Errorf("file not exists: %s", name)
	}
	if err != nil {
		return 0, err
	}
	return file.ID, nil
}

func (u *userRepo) CreateShareRecode(ctx context.Context, now time.Time, user *public.User, fids []int64, pwd, exp string) (string, error) {
	url_, err := url.Parse(public.OutterURL)
	if err != nil {
		u.log.Error(err)
		return "", err
	}
	var st sonyflake.Settings
	sf, err := sonyflake.New(st)
	if err != nil {
		return "", err
	}
	if id, err := sf.NextID(); err != nil {
		return "", err
	} else {
		url_ = url_.JoinPath(fmt.Sprintf("/share/%d", id))
		js, _ := json.Marshal(fids)
		share := public.ShareInfo{
			ID:   strconv.FormatUint(id, 10),
			Link: url_.String(), UserName: user.Name,
			Uid: user.ID, Fids: string(js), Created: now.Unix(),
		}
		_, err := u.engine.Insert(share)
		if err != nil {
			u.log.Error(err)
			return "", err
		}
		return url_.String(), nil
	}
}

func (u *userRepo) GetShares(ctx context.Context, id, pwd string) (string, []public.File, error) {
	shareInfo := public.ShareInfo{}
	u.engine.Where("id=?", id).Get(&shareInfo)
	if pwd != shareInfo.Password {
		return "", nil, fmt.Errorf("password of share incourrect")
	}
	var fids []int64
	json.Unmarshal([]byte(shareInfo.Fids), &fids)
	var finfos []public.File
	for _, fid := range fids {
		var finfo public.File
		u.engine.Where("id=?", fid).Get(&finfo)
		finfos = append(finfos, finfo)
	}
	return shareInfo.UserName, finfos, nil
}
