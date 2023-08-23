//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hifat/gock/internal/config"
	"github.com/hifat/gock/internal/database"
	"github.com/hifat/gock/internal/handler"
	"github.com/hifat/gock/internal/handler/helloHandler"
	"github.com/hifat/gock/internal/repository/helloRepository"
	"github.com/hifat/gock/internal/service/helloService"
)

var RepoSet = wire.NewSet(
	database.NewPostgresDBSet,
	helloRepository.NewHelloRepoSet,
)

var ServiceSet = wire.NewSet(
	helloService.NewHelloServiceSet,
)

var HandlerSet = wire.NewSet(
	handler.NewHandlerSet,
	helloHandler.NewHelloHandlerSet,
)

func InitializeAPI(config config.AppConfig) (Adapter, func()) {
	wire.Build(AdapterSet, RepoSet, ServiceSet, HandlerSet)
	return Adapter{}, nil
}
