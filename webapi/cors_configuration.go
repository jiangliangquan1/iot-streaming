package webapi

type CorsConfiguration struct {
	allowedOrigins        []string
	allowedOriginPatterns []string
	allowedMethods        []string
	allowedHeaders        []string
	exposeHeaders         []string
	allowedCredentials    bool
	maxAge                int64
}

func (c *CorsConfiguration) SetAllowedOrigins(origins []string) {
	c.allowedOrigins = origins
}

func (c *CorsConfiguration) GetAllowedOrigins() []string {
	return c.allowedOrigins
}

func (c *CorsConfiguration) SetAllowedOriginPatterns(patterns []string) {
	c.allowedOriginPatterns = patterns
}

func (c *CorsConfiguration) GetAllowedOriginPatterns() []string {
	return c.allowedOriginPatterns
}

func (c *CorsConfiguration) SetAllowedMethods(methods []string) {
	c.allowedMethods = methods
}

func (c *CorsConfiguration) GetAllowedMethods() []string {
	return c.allowedMethods
}

func (c *CorsConfiguration) SetAllowedHeaders(headers []string) {
	c.allowedHeaders = headers
}

func (c *CorsConfiguration) GetAllowedHeaders() []string {
	return c.allowedHeaders
}

func (c *CorsConfiguration) SetExposedHeaders(exposes []string) {
	c.exposeHeaders = exposes
}

func (c *CorsConfiguration) GetExposedHeaders() []string {
	return c.exposeHeaders
}

func (c *CorsConfiguration) SetAllowCredentials(allowCredentials bool) {
	c.allowedCredentials = allowCredentials
}

func (c *CorsConfiguration) GetAllowCredentials() bool {
	return c.allowedCredentials
}

func (c *CorsConfiguration) SetMaxAge(maxAge int64) {
	c.maxAge = maxAge
}

func (c *CorsConfiguration) GetMaxAge() int64 {
	return c.maxAge
}

func NewCorsConfiguration() *CorsConfiguration {
	return &CorsConfiguration{}
}
