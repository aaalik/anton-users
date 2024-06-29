package auth

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=type.go  -destination=auth_usecase_mock_test.go -package=auth
type AuthUsecase struct {
	log *logrus.Logger
	ur  iUserRepo
	jcu iJwtConfUtil
}

type iUserRepo interface {
	CreateTx(ctx context.Context) (*sqlx.Tx, error)
	RollbackTx(ctx context.Context, tx *sqlx.Tx) error
	CommitTx(ctx context.Context, tx *sqlx.Tx) error

	FetchUserLogin(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
}

type iJwtConfUtil interface {
	GetSecret() string
	GetSecretRefresh() string
	GetExpire() int
}
