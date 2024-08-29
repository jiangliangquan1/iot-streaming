package webapi

import (
	"github.com/bmatcuk/doublestar/v4"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type IWebApiServer interface {
	Run() error
	AddApiController(controller ApiController) error
	SetWebApiConfigure(c WebApiConfigurer) error
}

type WebApiServerOption struct {
	Port        int16
	UrlPathRoot string
	C           WebApiConfigurer
}

type WebApiServer struct {
	option *WebApiServerOption
	gin    *gin.Engine
	router *gin.RouterGroup
	logger *logrus.Logger

	controllers []ApiController

	configure WebApiConfigurer

	interceptorRegister *InterceptorRegister
	corsRegistry        *CorsRegistry
}

func (web *WebApiServer) SetWebApiConfigure(c WebApiConfigurer) error {
	if c == nil {
		panic("WebApiConfigurer is not allowed nil!")
	}

	web.configure = c

	web.interceptorRegister = &InterceptorRegister{}
	web.corsRegistry = &CorsRegistry{}

	c.AddCorsMappings(web.corsRegistry)

	c.AddInterceptors(web.interceptorRegister)

	return web.doSetWebApiConfigure()
}

func (web *WebApiServer) AddApiController(controller ApiController) error {

	web.controllers = append(web.controllers, controller)

	controller.RegisterRoute(web.router)

	return nil
}

func (web *WebApiServer) Run() error {
	web.gin.Run(":" + strconv.Itoa(int(web.option.Port)))

	return nil
}

func (web *WebApiServer) doSetWebApiConfigure() error {
	if web.corsRegistry != nil {
		web.doSetCors(web.corsRegistry)
	}

	if web.interceptorRegister != nil {
		web.doSetInterceptors(web.interceptorRegister)
	}

	return nil
}

func (web *WebApiServer) doSetCors(registry *CorsRegistry) error {
	corsMapping := registry.GetCorsConfigurations()
	if corsMapping == nil || len(corsMapping) == 0 {
		return nil
	}

	web.gin.Use(func(context *gin.Context) {
		for k, v := range corsMapping {
			match, _ := doublestar.Match(k, context.Request.URL.Path)
			if match {
				applyCORS(context, v)
				break
			}
		}

		context.Next()
	})

	return nil

}

func applyCORS(ctx *gin.Context, cfg *CorsConfiguration) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", strings.Join(cfg.GetAllowedOrigins(), ","))
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(cfg.GetAllowedMethods(), ","))
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(cfg.GetAllowedHeaders(), ","))
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", strings.Join(cfg.GetExposedHeaders(), ","))

	if cfg.GetAllowCredentials() {
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	if cfg.GetMaxAge() > 0 {
		ctx.Writer.Header().Set("Access-Control-Max-Age", strconv.FormatInt(cfg.GetMaxAge(), 10))
	}

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}
}

func (web *WebApiServer) doSetInterceptors(registry *InterceptorRegister) error {
	interceptors := registry.GetInterceptors()

	if len(interceptors) <= 0 {
		return nil
	}

	for _, v := range interceptors {
		web.gin.Use(func(context *gin.Context) {

			if matchInterceptorPathPattern(v.GetIncludePatterns(), v.GetExcludePatterns(), context.Request.URL.Path) {

				if !v.GetInterceptor().PreHandle(context) {
					context.Abort()
					return
				}

				context.Next()

				if !v.GetInterceptor().PostHandle(context) {
					context.Abort()
					return
				}

			} else {
				context.Next()
			}

		})
	}

	return nil

}

func matchInterceptorPathPattern(includes []string, excludes []string, reqUrl string) bool {
	for _, p := range excludes {
		if match, _ := doublestar.Match(p, reqUrl); match {
			return false //请求路径匹配上排除目录，无需执行拦截器
		}
	}

	for _, p := range includes {
		if match, _ := doublestar.Match(p, reqUrl); match {
			return true //请求路径匹配上包含目录，需执行拦截器
		}
	}

	return false
}

// NewWebApiServer 构造器
func NewWebApiServer(opt *WebApiServerOption, log *logrus.Logger, controllers []ApiController) IWebApiServer {
	g := gin.New()

	server := &WebApiServer{gin: g,
		option: &WebApiServerOption{Port: opt.Port, UrlPathRoot: opt.UrlPathRoot},
		logger: log}

	if opt.C != nil {
		server.SetWebApiConfigure(opt.C)
	}

	server.router = g.Group(opt.UrlPathRoot)

	for _, c := range controllers {
		server.AddApiController(c)
	}

	return server
}
