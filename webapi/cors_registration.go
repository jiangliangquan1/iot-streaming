package webapi

type CorsRegistration struct {
	pathPattern string
	config      *CorsConfiguration
}

func (c *CorsRegistration) AllowedOrigins(origins ...string) *CorsRegistration {
	c.config.SetAllowedOrigins(origins)
	return c
}

func (c *CorsRegistration) AllowedOriginPatterns(patterns ...string) *CorsRegistration {
	c.config.SetAllowedOriginPatterns(patterns)
	return c
}

func (c *CorsRegistration) AllowedMethods(methods ...string) *CorsRegistration {
	c.config.SetAllowedMethods(methods)
	return c
}

func (c *CorsRegistration) AllowedHeaders(headers ...string) *CorsRegistration {
	c.config.SetAllowedHeaders(headers)
	return c
}

func (c *CorsRegistration) ExposedHeaders(headers ...string) *CorsRegistration {
	c.config.SetExposedHeaders(headers)
	return c
}

func (c *CorsRegistration) AllowCredentials(allowCredentials bool) *CorsRegistration {
	c.config.SetAllowCredentials(allowCredentials)
	return c
}

func (c *CorsRegistration) MaxAge(maxAge int64) *CorsRegistration {
	c.config.SetMaxAge(maxAge)
	return c
}

func (c *CorsRegistration) getPathPattern() string {
	return c.pathPattern
}

func (c *CorsRegistration) getCorsConfiguration() *CorsConfiguration {
	return c.config
}

func NewCorsRegistration(pattern string) *CorsRegistration {
	return &CorsRegistration{pathPattern: pattern, config: &CorsConfiguration{}}
}
