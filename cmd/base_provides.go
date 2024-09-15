package main

import (
	userauth "github.com/jiangliangquan1/iot-streaming/authorization/user-auth"
	"github.com/jiangliangquan1/iot-streaming/basisdata/devices"
	"github.com/jiangliangquan1/iot-streaming/database"
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

func ProvideApiControllers(c1 *zlwebhook.ZlWebHookController, c2 *devices.DeviceController, c3 *userauth.UserController) []webapi.ApiController {

	var list []webapi.ApiController

	list = append(list, c1, c2, c3)

	return list

}

func ProvideDatabaseConnectOptions(v *viperex.ViperEx) *database.ConnectOptions {
	driver := v.GetString("datasource.driver")
	host := v.GetString("datasource.host")
	port := v.GetInt("datasource.port")
	dbName := v.GetString("datasource.dbname")
	username := v.GetString("datasource.username")
	password := v.GetString("datasource.password")
	charset := v.GetString("datasource.charset")
	loc := v.GetString("datasource.loc")

	return &database.ConnectOptions{
		Driver:   driver,
		Host:     host,
		Port:     int16(port),
		DBName:   dbName,
		Username: username,
		Password: password,
		Charset:  charset,
		Loc:      loc,
	}
}
