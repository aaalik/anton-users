package auth

import (
	"context"
	"net/http"

	"github.com/aaalik/anton-users/internal/service"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=type.go  -destination=user_handler_mock_test.go -package=user
type AuthHandler struct {
	httpRes iHttpResponse
	log     *logrus.Logger
	au      iAuthUC
}

type iAuthUC interface {
	Login(ctx context.Context, username, password string) (*service.ResponseLogin, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, int64, error)
	Register(ctx context.Context, request *service.RequestRegister) (*service.ResponseRegister, error)
}

type iHttpResponse interface {
	Nay(w http.ResponseWriter, r *http.Request, status int, err error)
	Yay(w http.ResponseWriter, r *http.Request, status int, content interface{})
	HTMLYay(w http.ResponseWriter, r *http.Request, status int, content string)
	DataYay(w http.ResponseWriter, r *http.Request, filename string, content []byte)
}
