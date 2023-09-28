package main

import (
	"context"
	"fmt"
	"net/http"
	"url_shortener/config"
	httpApi "url_shortener/internal/controller/http_api"
	"url_shortener/internal/repository"
	"url_shortener/internal/service"
	dbConn "url_shortener/pkg/db_conn"
	"url_shortener/pkg/db_conn/postgres"
	"url_shortener/pkg/logger"
	"url_shortener/pkg/logger/zerolog"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

func main() {
	ctx := context.Background()
	var cfg config.Config
	err := confita.NewLoader(
		env.NewBackend(),
	).Load(ctx, &cfg)
	var logger logger.ILogger = zerolog.NewLogger()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to parse config: %v", err))
	}
	var DBConnector dbConn.IDBConnector
	DBConnector, err = postgres.NewDBConnector(cfg.Database)
	defer DBConnector.CloseConnect()
	if err != nil {
		logger.Fatal(err.Error())
	}
	repo := repository.NewRepository(DBConnector, logger)
	service := service.NewService(logger, repo.GetLinkRepository())
	logger.Info(fmt.Sprint(cfg))
	server := httpApi.SetupRouter(logger, service)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), server)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
