//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/config"
	"github.com/hifat/gock/internal/database"
	"github.com/hifat/gock/internal/handler"
	"github.com/hifat/gock/internal/handler/taskHandler"
	"github.com/hifat/gock/internal/repository/taskRepository"
	"github.com/hifat/gock/internal/service/taskService"
)

var RepoSet = wire.NewSet(
	database.NewPostgresDBSet,
	taskRepository.NewTaskRepoSet,
)

var ServiceSet = wire.NewSet(
	taskService.NewTaskServiceSet,
)

var HandlerSet = wire.NewSet(
	handler.NewHandlerSet,
	taskHandler.NewTaskHandlerSet,
)

func InitializeAPI(config config.AppConfig) (Adapter, func()) {
	wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	return Adapter{}, nil
}
