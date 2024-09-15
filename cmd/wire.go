//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	userauth "github.com/jiangliangquan1/iot-streaming/authorization/user-auth"
	"github.com/jiangliangquan1/iot-streaming/basisdata/devices"
	"github.com/jiangliangquan1/iot-streaming/database"
	"github.com/jiangliangquan1/iot-streaming/logger"
	"github.com/jiangliangquan1/iot-streaming/repository"
	"github.com/jiangliangquan1/iot-streaming/viperex"
	"github.com/jiangliangquan1/iot-streaming/webapi"
	"github.com/jiangliangquan1/iot-streaming/zlmediaserver"
	"github.com/jiangliangquan1/iot-streaming/zlmediaserver/zlwebhook"
)

var ProviderSet = wire.NewSet(
	ProvideViperExOption,
	viperex.NewViperEx,
	ProvideLoggerConfigOption,
	logger.NewLogger,
	NewWebApiConfigurer,
	ProvideWebApiServerOption,
	zlwebhook.NewZlWebHookController,
	devices.NewDeviceController,
	webapi.NewWebApiServer,
	ProvideApiControllers,
	ProvideDatabaseConnectOptions,
	database.NewDataBase,
	repository.NewDeviceRepository,
	repository.NewUserRepository,
	userauth.NewUserService,
	userauth.NewUserController,
	userauth.NewJwtManager,
	userauth.NewUserAuthInterceptor,
	devices.NewDeviceService,
	repository.NewZlMediaServerRepository,
	zlmediaserver.NewListService,
	zlmediaserver.NewListController,
	NewIotStreamingApp,
)

func InitializeApp(configfile string) App {
	wire.Build(
		//wire.Value(configfile),
		ProviderSet)
	return nil
}
