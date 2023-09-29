package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"url_shortener/config"
	httpApi "url_shortener/internal/controller/http"
	"url_shortener/internal/repository"
	"url_shortener/internal/service"
	"url_shortener/pkg/db/postgres"
	"url_shortener/pkg/logger/zerolog"
	"url_shortener/pkg/server"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

func main() {
	ctx := context.Background()
	var cfg config.Config
	err := confita.NewLoader(
		env.NewBackend(),
	).Load(ctx, &cfg)
	logger := zerolog.NewLogger()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to parse config: %v", err))
	}
	db, err := postgres.NewDB(cfg.Database)
	defer db.CloseConnect()
	if err != nil {
		logger.Fatal(err.Error())
	}
	repo := repository.NewRepository(db, logger)
	service := service.NewService(logger, repo.GetLinkRepository())
	logger.Info(fmt.Sprint(cfg))
	router := httpApi.SetupRouter(logger, service)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg.Port, router); err != nil {
			logger.Fatal(err.Error())
		}
	}()
	logger.Info("app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Info("app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}
