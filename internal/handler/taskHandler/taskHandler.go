package taskHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
)

var NewTaskHandlerSet = wire.NewSet(NewTaskHandler)

type TaskHandler struct {
	taskSrv taskDomain.TaskService
}

func NewTaskHandler(taskSrv taskDomain.TaskService) TaskHandler {
	return TaskHandler{taskSrv}
}

func (h *TaskHandler) Get(ctx *gin.Context) {
	res := h.taskSrv.Get()
	ctx.JSON(200, gin.H{
		"msg": res,
	})
}
