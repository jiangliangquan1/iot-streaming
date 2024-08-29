package webapi

import "github.com/gin-gonic/gin"

type Interceptor interface {
	PreHandle(c *gin.Context) bool
	PostHandle(c *gin.Context) bool
}
