package helloHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/helloDomain"
)

var NewHelloHandlerSet = wire.NewSet(NewHelloHandler)

type HelloHandler struct {
	helloSrv helloDomain.HelloService
}

func NewHelloHandler(helloSrv helloDomain.HelloService) HelloHandler {
	return HelloHandler{helloSrv}
}

func (h *HelloHandler) GetHello(ctx *gin.Context) {
	res := h.helloSrv.GetHello()
	ctx.JSON(200, gin.H{
		"msg": res,
	})
}
