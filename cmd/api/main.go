package main

import (
	"net/http"

	"github.com/aaalik/anton-users/cmd"
	"github.com/aaalik/anton-users/cmd/api/routes"
	"github.com/aaalik/anton-users/cmd/config"

	authHandler "github.com/aaalik/anton-users/internal/handler/httphandler/auth"
	userHandler "github.com/aaalik/anton-users/internal/handler/httphandler/user"
	userRepo "github.com/aaalik/anton-users/internal/repository/postgres/user"
	authUsecase "github.com/aaalik/anton-users/internal/usecase/auth"
	userUsecase "github.com/aaalik/anton-users/internal/usecase/user"
	"github.com/aaalik/anton-users/pkg/database"
	"github.com/aaalik/anton-users/pkg/hasher"
	httpResponse "github.com/aaalik/anton-users/pkg/httpresponse"
	"github.com/aaalik/anton-users/pkg/jwtconf"
	"github.com/aaalik/anton-users/pkg/jwtgen"
	"github.com/aaalik/anton-users/pkg/utils"
)

func main() {
	// init config & infastructures
	cf := config.NewConfig()
	logr := cf.NewLogrus()
	dbr, dbw := cf.NewPostgres()
	httpRes := httpResponse.NewHttpResponse(logr)
	jwtConf := jwtconf.NewJwtConf(cf.JwtConf.Secret, cf.JwtConf.SecretRefresh, cf.JwtConf.Expire)
	dbUtils := database.NewDB(dbr, dbw)
	randUtils := utils.NewRandomUtils()
	hashUtils := hasher.NewHasherUtils()
	jwtUtils := jwtgen.NewJwtUtils()

	srv := cmd.NewServer(cf, logr, dbr, dbw)

	// init repositories
	logr.Infoln("Initializing repositories")
	userRP := userRepo.New(logr, dbr, dbw)

	// init usecase
	logr.Infoln("Initializing usecases")
	authUC := authUsecase.New(userRP, jwtConf, dbUtils, randUtils, hashUtils, jwtUtils)
	userUC := userUsecase.New(userRP, dbUtils, randUtils, hashUtils)

	// init handler
	logr.Infoln("Initializing handlers")
	routerHandler := &routes.RouteHandler{
		AuthHandler: authHandler.New(httpRes, logr, authUC),
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
