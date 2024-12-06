package server

import (
	"context"
	v1 "driver-back/api/user/v1"
	"driver-back/internal/conf"
	"driver-back/internal/service"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userSvc *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)))
	opts = append(opts, http.Middleware(
		selector.Server(
			// jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			// 	fmt.Printf("+++++++++++++++++++++++++++++++")
			// 	if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// 	}
			// 	return []byte(service.SecretKey), nil
			// },
			// 	jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			// ),

			jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(service.SecretKey), nil
			}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256), jwt.WithClaims(func() jwtv5.Claims {
				return &jwtv5.MapClaims{}
			})),
		).
			Match(NewWhiteListMatcher()).
			Build(),
	))
	srv := http.NewServer(opts...)
	v1.RegisterUserSvcHTTPServer(srv, userSvc)
	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/user.v1.UserSvc/Login"] = struct{}{}
	whiteList["/user.v1.UserSvc/Register"] = struct{}{}
	whiteList["/user.v1.UserSvc/GetShare"] = struct{}{}
	whiteList["/user.v1.UserSvc/DownloadShare"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		fmt.Println(operation)
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
