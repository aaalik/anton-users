package auth

import (
	"github.com/sirupsen/logrus"
)

func New(log *logrus.Logger, userRP iUserRepo, jcu iJwtConfUtil) *AuthUsecase {
	return &AuthUsecase{
		log: log,
		ur:  userRP,
		jcu: jcu,
	}
}
