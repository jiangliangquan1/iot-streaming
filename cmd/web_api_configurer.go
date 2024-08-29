package main

import (
	"github.com/jiangliangquan1/iot-streaming/webapi"
	"github.com/sirupsen/logrus"
)

type WebApiConfigurer struct {
}

func (w *WebApiConfigurer) AddInterceptors(registry *webapi.InterceptorRegister) {

}

func (w *WebApiConfigurer) AddCorsMappings(registry *webapi.CorsRegistry) {
	registry.AddMapping("/**").
		AllowedHeaders("*").
		AllowCredentials(true).
		AllowedMethods("POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE").
		ExposedHeaders("*").
		AllowedOriginPatterns("*")
}

func NewWebApiConfigurer(log *logrus.Logger) webapi.WebApiConfigurer {
	return &WebApiConfigurer{}
}
