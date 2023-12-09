package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task-manager/config"
	"task-manager/internal/controller/http"
	"task-manager/internal/repository"
	"task-manager/internal/service"
	"task-manager/pkg/database"
	"task-manager/pkg/server"
)

func Run(cfg *config.Config) {
	pg, err := database.New(context.Background(), cfg.ConnURI)
	if err != nil {
		log.Println(err)
	}

	repos := repository.NewRepositories(pg)
	deps := &service.Dependencies{Repos: repos}
	services := service.NewServices(deps)

	rout := mux.NewRouter()
	http.NewRoutes(services, rout)
	srv := server.NewServer(rout, server.Port(cfg.ServerCfg.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err = srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
