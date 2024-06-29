package auth

import (
	"github.com/sirupsen/logrus"
)

func New(httpRes iHttpResponse, log *logrus.Logger, authUC iAuthUC) *AuthHandler {
	return &AuthHandler{
		httpRes: httpRes,
		log:     log,
		au:      authUC,
	}
}
