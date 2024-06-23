package user

import (
	"github.com/sirupsen/logrus"
)

func New(httpRes iHttpResponse, log *logrus.Logger, userUC iUserUC) *UserHandler {
	return &UserHandler{
		httpRes: httpRes,
		log:     log,
		uu:      userUC,
	}
}
