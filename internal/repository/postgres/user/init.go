package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func New(log *logrus.Logger, dbr, dbw *sqlx.DB) *UserRepository {
	return &UserRepository{
		log: log,
		dbr: dbr,
		dbw: dbw,
	}
}
