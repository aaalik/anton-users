package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=type.go  -destination=user_repo_mock_test.go -package=user
type UserRepository struct {
	log *logrus.Logger
	dbr *sqlx.DB
	dbw *sqlx.DB
}
