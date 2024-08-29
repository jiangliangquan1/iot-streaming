package webapi

type WebApiConfigurer interface {
	AddInterceptors(registry *InterceptorRegister)
	AddCorsMappings(registry *CorsRegistry)
}
