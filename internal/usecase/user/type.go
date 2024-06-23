package user

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=type.go  -destination=user_usecase_mock_test.go -package=user
type UserUsecase struct {
	log *logrus.Logger
	ur  iUserRepo
}

type iUserRepo interface {
	CreateTx(ctx context.Context) (*sqlx.Tx, error)
	RollbackTx(ctx context.Context, tx *sqlx.Tx) error
	CommitTx(ctx context.Context, tx *sqlx.Tx) error

	CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
	UpdateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
	DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error
	DetailUser(ctx context.Context, id string) (*model.User, error)
	ListUser(ctx context.Context, request *service.RequestListUser) ([]*model.User, error)
	CountUsers(ctx context.Context, request *service.RequestListUser) (int32, error)
}
