package taskHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/hifat/gock/internal/domain/taskDomain"
	"github.com/hifat/gock/internal/handler/handlerResponse"
	"github.com/hifat/gock/internal/utils/uResponse"
)

var NewTaskHandlerSet = wire.NewSet(NewTaskHandler)

type TaskHandler struct {
	taskSrv taskDomain.TaskService
}

func NewTaskHandler(taskSrv taskDomain.TaskService) TaskHandler {
	return TaskHandler{taskSrv}
}

func (h *TaskHandler) Get(c *gin.Context) {
	res := []taskDomain.Task{}
	if err := h.taskSrv.Get(c.Request.Context(), &res); err != nil {
		handlerResponse.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, uResponse.SuccessResponse{
		Items: res,
	})
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		handlerResponse.BadRequest(c, err)
		return
	}

	var res taskDomain.Task
	if err := h.taskSrv.GetByID(c.Request.Context(), &res, taskID); err != nil {
		handlerResponse.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req taskDomain.TaskRequest
	if err := c.ShouldBind(&req); err != nil {
		handlerResponse.BadRequest(c, err)
		return
	}

	res, err := h.taskSrv.Create(c.Request.Context(), &req)
	if err != nil {
		handlerResponse.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h *TaskHandler) Update(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		handlerResponse.BadRequest(c, err)
		return
	}

	var req taskDomain.TaskRequest
	if err := c.ShouldBind(&req); err != nil {
		handlerResponse.BadRequest(c, err)
		return
	}

	res, err := h.taskSrv.Update(c.Request.Context(), &req, taskID)
	if err != nil {
		handlerResponse.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, uResponse.SuccessResponse{
		Item: res,
	})
}

func (h *TaskHandler) Delete(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		handlerResponse.BadRequest(c, err)
		return
	}

	if err := h.taskSrv.Delete(c.Request.Context(), taskID); err != nil {
		handlerResponse.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, uResponse.SuccessResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
