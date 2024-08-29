package webapi

type CorsRegistry struct {
	corsRegistrations []*CorsRegistration
}

func (c *CorsRegistry) AddMapping(pathPattern string) *CorsRegistration {

	registration := &CorsRegistration{pathPattern: pathPattern, config: &CorsConfiguration{}}

	c.corsRegistrations = append(c.corsRegistrations, registration)

	return registration
}

func (c *CorsRegistry) GetCorsConfigurations() map[string]*CorsConfiguration {
	result := make(map[string]*CorsConfiguration)

	for _, r := range c.corsRegistrations {
		result[r.pathPattern] = r.config
	}

	return result
}
