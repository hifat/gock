package handlerResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/gock/internal/utils/ernos"
)

func Error(ctx *gin.Context, err error) {
	if errRes, ok := err.(ernos.Ernos); ok {
		ctx.AbortWithStatusJSON(errRes.Status, errRes)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, ernos.Ernos{
		Message: http.StatusText(http.StatusInternalServerError),
	})
}

func BadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, ernos.Ernos{
		Message: err.Error(),
	})
}
