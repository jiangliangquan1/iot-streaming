package zlwebhook

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ZlWebHookController struct {
	logger *logrus.Logger
}

func (z *ZlWebHookController) RegisterRoute(group *gin.RouterGroup) {
	group.GET("test", z.test)
}

func (z *ZlWebHookController) test(ctx *gin.Context) {
	ctx.JSON(200, "111")
}

func NewZlWebHookController(l *logrus.Logger) *ZlWebHookController {
	return &ZlWebHookController{logger: l}
}
