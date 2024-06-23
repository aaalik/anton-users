package user

import (
	"github.com/sirupsen/logrus"
)

func New(log *logrus.Logger, userRP iUserRepo) *UserUsecase {
	return &UserUsecase{
		log: log,
		ur:  userRP,
	}
}
