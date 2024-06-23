package config

import (
	"github.com/sirupsen/logrus"
)

func (cf Config) NewLogrus() *logrus.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()

	return log
}
