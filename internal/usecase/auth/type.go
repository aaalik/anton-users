package auth

import (
	"context"

	"github.com/aaalik/anton-users/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=type.go  -destination=auth_usecase_mock_test.go -package=auth
type AuthUsecase struct {
	ur  iUserRepo
	jcu iJwtConfUtils
	dbu iDatabaseUtils
	ru  iRandomUtils
	hu  iHasherUtils
	jwu iJwtUtils
}

type iUserRepo interface {
	FetchUserLogin(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, tx *sqlx.Tx, user *model.User) error
}

type iJwtConfUtils interface {
	GetSecret() string
	GetSecretRefresh() string
	GetExpire() int
}

type iDatabaseUtils interface {
	ExecuteTx(ctx context.Context, tx *sqlx.Tx, fn func(ctx context.Context, tx *sqlx.Tx) error) error
}

type iRandomUtils interface {
	UniqueID() string
}

type iHasherUtils interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type iJwtUtils interface {
	GenerateToken(claims jwt.MapClaims, secretKey string) (string, error)
	ParseToken(tokenString, secretKey string) (*jwt.Token, error)
}
