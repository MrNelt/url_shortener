package httpApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() http.Handler {
	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})

	// place openapi.json
	router.StaticFile("/swagger/api.json", "api/openapi.json")
	router.Static("/swagger-ui", "static/swagger-ui/dist")
	return router
}
