//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jiangliangquan1/iot-streaming/basisdata/devices"
	"github.com/jiangliangquan1/iot-streaming/logger"
	"github.com/jiangliangquan1/iot-streaming/viperex"
	"github.com/jiangliangquan1/iot-streaming/webapi"
	"github.com/jiangliangquan1/iot-streaming/zlwebhook"
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
	NewIotStreamingApp,
)

func InitializeApp(configfile string) App {
	wire.Build(
		//wire.Value(configfile),
		ProviderSet)
	return nil
}
