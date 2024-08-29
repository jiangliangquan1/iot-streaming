package devices

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DeviceController struct {
}

func (d *DeviceController) RegisterRoute(group *gin.RouterGroup) {

}

func NewDeviceController(l *logrus.Logger) *DeviceController {
	return &DeviceController{}
}
