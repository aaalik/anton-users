package user

import (
	"context"
	"net/http"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=type.go  -destination=user_handler_mock_test.go -package=user
type UserHandler struct {
	httpRes iHttpResponse
	log     *logrus.Logger
	uu      iUserUC
}

type iUserUC interface {
	CreateUser(ctx context.Context, request *service.RequestCreateUser) (*model.User, error)
	UpdateUser(ctx context.Context, request *service.RequestUpdateUser) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
	DetailUser(ctx context.Context, id string) (*model.User, error)
	ListUser(ctx context.Context, request *service.RequestListUser) ([]*model.User, int32, error)
}

type iHttpResponse interface {
	Nay(w http.ResponseWriter, r *http.Request, status int, err error)
	Yay(w http.ResponseWriter, r *http.Request, status int, content interface{})
	HTMLYay(w http.ResponseWriter, r *http.Request, status int, content string)
	DataYay(w http.ResponseWriter, r *http.Request, filename string, content []byte)
}
