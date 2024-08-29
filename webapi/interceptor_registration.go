package webapi

type InterceptorRegistration struct {
	interceptor     Interceptor
	includePatterns []string
	excludePatterns []string
}

func (i *InterceptorRegistration) AddPathPatterns(paths ...string) *InterceptorRegistration {
	i.AddPathPatternList(paths)
	return i
}

func (i *InterceptorRegistration) AddPathPatternList(paths []string) *InterceptorRegistration {
	i.includePatterns = append(i.includePatterns, paths...)
	return i
}

func (i *InterceptorRegistration) ExcludePathPatterns(paths ...string) *InterceptorRegistration {
	i.ExcludePathPatternList(paths)
	return i
}

func (i *InterceptorRegistration) ExcludePathPatternList(paths []string) *InterceptorRegistration {
	i.excludePatterns = append(i.excludePatterns, paths...)
	return i
}

func (i *InterceptorRegistration) GetInterceptor() Interceptor {
	return i.interceptor
}

func (i *InterceptorRegistration) GetIncludePatterns() []string {
	return i.includePatterns
}

func (i *InterceptorRegistration) GetExcludePatterns() []string {
	return i.excludePatterns
}

func NewInterceptorRegistration(i Interceptor) *InterceptorRegistration {
	return &InterceptorRegistration{interceptor: i}
}
