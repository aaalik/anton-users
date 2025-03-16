package user

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=type.go  -destination=user_usecase_mock_test.go -package=user
type UserUsecase struct {
	ur  iUserRepo
	dbu iDatabaseUtils
	ru  iRandomUtils
	hu  iHasherUtils
}

type iUserRepo interface {
	CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
	UpdateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
	DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error
	DetailUser(ctx context.Context, id string) (*model.User, error)
	ListUser(ctx context.Context, request *service.RequestListUser) ([]*model.User, error)
	CountUsers(ctx context.Context, request *service.RequestListUser) (int32, error)
}

type iDatabaseUtils interface {
	ExecuteTx(ctx context.Context, tx *sqlx.Tx, fn func(ctx context.Context, tx *sqlx.Tx) error) error
}

type iRandomUtils interface {
	UniqueID() string
}

type iHasherUtils interface {
	HashPassword(password string) (string, error)
}
