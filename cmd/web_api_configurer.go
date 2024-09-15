package main

import (
	userauth "github.com/jiangliangquan1/iot-streaming/authorization/user-auth"
	"github.com/jiangliangquan1/iot-streaming/webapi"
	"github.com/sirupsen/logrus"
)

type WebApiConfigurer struct {
	userAuthInterceptor *userauth.UserAuthInterceptor
}

func (w *WebApiConfigurer) AddInterceptors(registry *webapi.InterceptorRegister) {
	registry.AddInterceptor(w.userAuthInterceptor).AddPathPatterns("/**").
		ExcludePathPatterns("/iot-streaming/api/v1/users/sign-up", "/iot-streaming/api/v1/users/login")
}

func (w *WebApiConfigurer) AddCorsMappings(registry *webapi.CorsRegistry) {
	registry.AddMapping("/**").
		AllowedHeaders("*").
		AllowCredentials(true).
		AllowedMethods("POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE").
		ExposedHeaders("*").
		AllowedOriginPatterns("*")
}

func NewWebApiConfigurer(log *logrus.Logger, userAuthInterceptor *userauth.UserAuthInterceptor) webapi.WebApiConfigurer {
	return &WebApiConfigurer{userAuthInterceptor: userAuthInterceptor}
}
