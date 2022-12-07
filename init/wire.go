//go:build wireinject
// +build wireinject

package main

import (
	"github.com/MaxWxaM/linebot/internal/app"
	"github.com/MaxWxaM/linebot/internal/app/http/controller"

	"github.com/MaxWxaM/linebot/internal/app/config"
	"github.com/MaxWxaM/linebot/internal/app/http/route"
	"github.com/MaxWxaM/linebot/internal/app/repo"
	"github.com/MaxWxaM/linebot/internal/app/service"
	"github.com/google/wire"
)

var (
	controllerProviderSet = wire.NewSet(
		controller.NewMsgController,
	)
	serviceProviderSet = wire.NewSet(
		service.NewMsgService,
	)

	repositoryProviderSet = wire.NewSet(
		repo.NewMongoMsgRepo,
	)

	linebotProviderSet = wire.NewSet(
		config.NewLineBotClientConfig,
		config.NewLineBotClient,
	)
	mongoProviderSet = wire.NewSet(
		config.NewMongoClientConfig,
		config.NewMongoClient,
	)
	routeProviderSet = wire.NewSet(
		route.NewRoute,
	)
	serverProviderSet = wire.NewSet(
		config.NewHttpOptions,
		app.NewServer,
	)
)

var providerSet = wire.NewSet(
	controllerProviderSet,
	routeProviderSet,
	linebotProviderSet,
	serviceProviderSet,
	repositoryProviderSet,
	mongoProviderSet,
	serverProviderSet,
)

func InitApp() (*app.Server, func(), error) {
	panic(wire.Build(providerSet))
}
