package webapi

type InterceptorRegister struct {
	interceptors []*InterceptorRegistration
}

func (i *InterceptorRegister) AddInterceptor(interceptor Interceptor) *InterceptorRegistration {
	registration := NewInterceptorRegistration(interceptor)
	i.interceptors = append(i.interceptors, registration)
	return registration
}

func (i *InterceptorRegister) GetInterceptors() []*InterceptorRegistration {
	return i.interceptors
}

func NewInterceptorRegister() *InterceptorRegister {
	return &InterceptorRegister{}
}
