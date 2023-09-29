package httpApi

import (
	"net/http"
	"url_shortener/internal/controller/http/link"
	"url_shortener/internal/service"
	"url_shortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(logger logger.ILogger, service service.IService) http.Handler {
	router := gin.Default()
	var handler link.IHandler = link.NewHandler(logger, service)
	router.POST("/make_shorter", handler.MakeShorter)
	router.GET("/info/:id", handler.Info)
	router.GET("/:short_suffix", handler.Redirect)
	router.StaticFile("/swagger/api.json", "api/openapi.json")
	router.Static("/swagger-ui", "static/swagger-ui/dist")
	return router
}
