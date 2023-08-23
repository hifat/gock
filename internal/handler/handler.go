package handler

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/handler/taskHandler"
)

var NewHandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	TaskHandler taskHandler.TaskHandler
}

func NewHandler(TaskHandler taskHandler.TaskHandler) Handler {
	return Handler{
		TaskHandler,
	}
}
