package main

import (
	"github.com/jiangliangquan1/iot-streaming/basisdata/devices"
	"github.com/jiangliangquan1/iot-streaming/logger"
	"github.com/jiangliangquan1/iot-streaming/viperex"
	"github.com/jiangliangquan1/iot-streaming/webapi"
	"github.com/jiangliangquan1/iot-streaming/zlwebhook"
)

func ProvideViperExOption(configfile string) *viperex.Option {
	return &viperex.Option{ConfigFile: configfile, BindEnv: true}
}

func ProvideLoggerConfigOption(vieperex *viperex.ViperEx) *logger.ConfigOptions {

	level := vieperex.GetString("logrus.level")
	filename := vieperex.GetString("logrus.filename")
	maxSize := vieperex.GetInt("logrus.maxSize")
	maxBackups := vieperex.GetInt("logrus.maxBackups")
	maxAge := vieperex.GetInt("logrus.maxAge")

	return &logger.ConfigOptions{
		Level:      level,
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
}

func ProvideWebApiServerOption(v *viperex.ViperEx, c webapi.WebApiConfigurer) *webapi.WebApiServerOption {

	urlPathRoot := v.GetString("webapi.url-path-root")
	port := v.GetInt("webapi.port")

	return &webapi.WebApiServerOption{UrlPathRoot: urlPathRoot, Port: int16(port), C: c}
}

type ControllerAddNone struct {
}

func ProvideApiControllers(c1 *zlwebhook.ZlWebHookController, c2 *devices.DeviceController) []webapi.ApiController {

	var list []webapi.ApiController

	list = append(list, c1, c2)

	return list

}
