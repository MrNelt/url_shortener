package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"url_shortener/config"
	httpApi "url_shortener/internal/controller/http_api"
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
	logger.Info(fmt.Sprint(cfg))
	server := httpApi.SetupRouter()
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), server)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}
