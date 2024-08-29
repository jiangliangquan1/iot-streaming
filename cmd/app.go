package main

import (
	"fmt"
	"github.com/jiangliangquan1/iot-streaming/webapi"
)

type App interface {
	Run() error
}

type IotStreamingApp struct {
	_webApiServer webapi.IWebApiServer
}

func (app *IotStreamingApp) Run() error {
	fmt.Println("IotStreamingApp start run!")
	app._webApiServer.Run()
	return nil
}

func NewIotStreamingApp(webApiServer webapi.IWebApiServer) App {
	return &IotStreamingApp{_webApiServer: webApiServer}
}
