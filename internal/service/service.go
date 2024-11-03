package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

const SecretKey = "87485hji&*^&(YHLRFJDS)"
