package routes

import (
	"github.com/sirupsen/logrus"

	"github.com/aaalik/anton-users/internal/handler/httphandler/user"
)

type Route struct {
	log          *logrus.Logger
	routeHandler *RouteHandler
}

type RouteHandler struct {
	log         *logrus.Logger
	UserHandler *user.UserHandler
}
