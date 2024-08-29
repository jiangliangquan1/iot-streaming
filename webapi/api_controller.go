package webapi

import (
	"github.com/gin-gonic/gin"
)

type ApiController interface {
	RegisterRoute(group *gin.RouterGroup)
}
