package link

import (
	"errors"
	"net/http"
	"url_shortener/internal/dtos"
	errorsApi "url_shortener/internal/errors"
	"url_shortener/internal/service"
	"url_shortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	MakeShorter(ctx *gin.Context)
	Info(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}

type Handler struct {
	logger  logger.ILogger
	service service.IService
}

func NewHandler(logger logger.ILogger, service service.IService) *Handler {
	return &Handler{logger: logger, service: service}
}

func (h *Handler) MakeShorter(ctx *gin.Context) {
	var linkDTO dtos.LinkDTO
	if err := ctx.ShouldBindJSON(&linkDTO); err != nil {
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}
	ID, err := h.service.SaveLink(linkDTO)
	switch {
	case errors.Is(err, errorsApi.ErrSuffixAlreadyExists):
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	case err != nil:
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"short_suffix": linkDTO.ShortSuffix,
		"id":           ID,
	})
}

func (h *Handler) Info(ctx *gin.Context) {
	ID := ctx.Param("id")
	link, err := h.service.GetLinkInfoFromID(ID)
	switch {
	case errors.Is(err, errorsApi.ErrIDNotFound):
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	case err != nil:
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}
	ctx.JSON(http.StatusOK, link)
}

func (h *Handler) Redirect(ctx *gin.Context) {
	url, err := h.service.GetUrlForRedirect(ctx.Param("short_suffix"))
	switch {
	case errors.Is(err, errorsApi.ErrSuffixNotFound):
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	case err != nil:
		h.logger.Error(err.Error())
		errorsApi.HandleError(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
