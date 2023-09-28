package errorsApi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ErrSuffixAlreadyExists = fmt.Errorf("suffix already exists")
	ErrSuffixNotFound      = fmt.Errorf("suffix not found")
	ErrLinkNotFound        = fmt.Errorf("link not found")
	ErrIDNotFound          = fmt.Errorf("ID not found")
	ErrTTL                 = fmt.Errorf("TTL_Unit between (SECONDS, MINUTES, HOURS, DAYS), TTL_Count >= 0")
)

func HandleError(ctx *gin.Context, status int, errMsg string, err error) {
	response := gin.H{
		"error": errMsg,
	}
	ctx.JSON(status, response)
}
