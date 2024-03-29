package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/gock/internal/handler"
)

type Route struct {
	router  *gin.RouterGroup
	handler handler.Handler
}

func New(router *gin.RouterGroup, h handler.Handler) *Route {
	return &Route{
		router:  router,
		handler: h,
	}
}

func (r *Route) Register() {
	v1 := r.router.Group("v1/api")

	tasks := v1.Group("/tasks")
	{
		tasks.GET("", r.handler.TaskHandler.Get)
		tasks.GET("/:taskID", r.handler.TaskHandler.GetByID)
		tasks.POST("", r.handler.TaskHandler.Create)
		tasks.PATCH("/:taskID", r.handler.TaskHandler.Update)
		tasks.DELETE("/:taskID", r.handler.TaskHandler.Delete)
	}
}
