package main

import (
	"net/http"

	"github.com/aaalik/anton-users/cmd"
	"github.com/aaalik/anton-users/cmd/api/routes"
	"github.com/aaalik/anton-users/cmd/config"

	userHandler "github.com/aaalik/anton-users/internal/handler/httphandler/user"
	userRepo "github.com/aaalik/anton-users/internal/repository/postgres/user"
	userUsecase "github.com/aaalik/anton-users/internal/usecase/user"
	httpResponse "github.com/aaalik/anton-users/pkg/httpresponse"
)

func main() {
	// init config & infastructures
	cf := config.NewConfig()
	logr := cf.NewLogrus()
	dbr, dbw := cf.NewPostgres()
	httpRes := httpResponse.NewHttpResponse(logr)

	srv := cmd.NewServer(cf, logr, dbr, dbw)

	// init repositories
	logr.Infoln("Initializing repositories")
	userRP := userRepo.New(logr, dbr, dbw)

	// init usecase
	logr.Infoln("Initializing usecases")
	userUC := userUsecase.New(logr, userRP)

	// init handler
	logr.Infoln("Initializing handlers")
	routerHandler := &routes.RouteHandler{
		UserHandler: userHandler.New(httpRes, logr, userUC),
	}

	// init routes
	r := routes.New(logr, routerHandler).Init()

	server := http.Server{
		Addr:    cf.Host.Address,
		Handler: r,
	}

	serverErr := make(chan error, 1)
	go func() {
		logr.Infof("User API serving at %s", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	// Graceful Stop handle
	srv.StopGracefully(&server, serverErr)
}
