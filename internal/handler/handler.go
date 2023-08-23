package handler

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/handler/helloHandler"
)

var NewHandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	HelloHandler helloHandler.HelloHandler
}

func NewHandler(HelloHandler helloHandler.HelloHandler) Handler {
	return Handler{
		HelloHandler,
	}
}
