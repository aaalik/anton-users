package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/aaalik/anton-users/cmd/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cf  config.Config
	log *logrus.Logger
	dbr *sqlx.DB
	dbw *sqlx.DB
}

func NewServer(
	cf config.Config,
	log *logrus.Logger,
	dbr *sqlx.DB,
	dbw *sqlx.DB,
) *Server {
	return &Server{
		cf:  cf,
		log: log,
		dbr: dbr,
		dbw: dbw,
	}
}

func (srv *Server) StopGracefully(server *http.Server, serverErr chan error) {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-shutdownChannel:
		srv.log.Info("Caught signal ", sig, " Stop Gracefully")

		timeoutCfg := srv.cf.StopTimeout
		timeout := time.Duration(timeoutCfg) * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		defer srv.CloseConnection()
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
		}
	case err := <-serverErr:
		if err != nil {
			log.Fatalf("server: %v", err)
		}
	}
}

func (srv *Server) CloseConnection() {
	srv.log.Println("closing connections...")

	wg := sync.WaitGroup{}
	if srv.dbr != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv.dbr.Close()
		}()
	}
	if srv.dbw != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv.dbw.Close()
		}()
	}
	if srv.log != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			srv.log.Writer().Close()
		}()
	}

	wg.Wait()
}
