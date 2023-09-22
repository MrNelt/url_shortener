package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"url_shortener/config"
	httpApi "url_shortener/internal/controller/http_api"
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
		os.Exit(1)
	}
	var DBConnector dbConn.IDBConnector
	DBConnector, err = postgres.NewDBConnector(cfg.Database)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	logger.Info(fmt.Sprint(cfg))
	server := httpApi.SetupRouter()
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), server)
	if err != nil {
		logger.Fatal(err.Error())
		err = DBConnector.CloseConnect()
		logger.Error(err.Error())
		os.Exit(1)
	}
	DBConnector.CloseConnect()
}
