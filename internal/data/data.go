package data

import (
	"driver-back/internal/conf"
	"driver-back/public"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Driver      string
	DBAddr      string
	MinioAccess MinioAccess
	Env         Env
}

type MinioAccess struct {
	KeyID    string
	KeyPass  string
	Endpoint string
}

type Env struct {
	StoreType string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	if c.Env.StoreType != "" {
		public.StoreType = c.Env.StoreType
	}
	d := &Data{
		DBAddr: c.Database.Source, Driver: c.Database.Driver,
		MinioAccess: MinioAccess{
			KeyID:    c.Minio.KeyID,
			KeyPass:  c.Minio.KeyPass,
			Endpoint: c.Minio.Endpoint,
		},
	}
	return d, cleanup, nil
}
