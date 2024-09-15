package devices

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DeviceController struct {
	logger        *logrus.Logger
	deviceService *DeviceService
}

func (d *DeviceController) RegisterRoute(group *gin.RouterGroup) {
	g := group.Group("v1/devices")

	g.POST("", d.create)
	g.GET(":id", d.getByID)
	g.PUT(":id", d.update)
	g.DELETE(":id", d.deleteByID)
}

func (d *DeviceController) create(ctx *gin.Context) {

	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var req DeviceCreateRequest
	ctx.BindJSON(&req)

	response, err := d.deviceService.Create(&req, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: response})
}

func (d *DeviceController) getByID(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("设备id参数：%s 错误", idStr), Result: nil})
		return
	}

	response, err := d.deviceService.GetByID(id, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: response})
}

func (d *DeviceController) update(ctx *gin.Context) {

	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("设备id参数：%s 错误", idStr), Result: nil})
		return
	}

	var req DeviceUpdateRequest
	ctx.BindJSON(&req)
	req.ID = id
	response, err := d.deviceService.Update(&req, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: response})
}

func (d *DeviceController) deleteByID(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("设备id参数：%s 错误", idStr), Result: nil})
		return
	}

	response, err := d.deviceService.DeleteByID(id, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: response})
}

func NewDeviceController(l *logrus.Logger, deviceService *DeviceService) *DeviceController {
	return &DeviceController{logger: l, deviceService: deviceService}
}
