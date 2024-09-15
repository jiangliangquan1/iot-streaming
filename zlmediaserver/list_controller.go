package zlmediaserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"net/http"
	"strconv"
)

// zlmediaserver 列表管理接口
type ListController struct {
	listService *ListService
}

func (l *ListController) RegisterRoute(group *gin.RouterGroup) {
	g := group.Group("v1/zlmediaserver")

	g.POST("", l.create)
	g.PUT("", l.update)
	g.GET(":id", l.getByID)
	g.GET("", l.getByServerID)
	g.DELETE(":id", l.deleteByID)
}

func (l *ListController) create(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var req CreateRequest
	ctx.BindJSON(&req)

	r, err := l.listService.Create(&req, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: r})
}

func (l *ListController) update(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var req UpdateRequest
	ctx.BindJSON(&req)

	r, err := l.listService.Update(&req, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: r})
}

func (l *ListController) getByID(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("id参数：%s 错误", idStr), Result: nil})
		return
	}

	r, err := l.listService.GetByID(id, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: r})
}

func (l *ListController) getByServerID(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	serverID := ctx.Query("server_id")
	if len(serverID) <= 0 {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("缺失server_id参数"), Result: nil})
		return
	}

	r, err := l.listService.GetByServerID(serverID, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: r})
}

func (l *ListController) deleteByID(ctx *gin.Context) {
	userContext, exist := ctx.Get("userContext")
	if !exist {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: fmt.Sprintf("id参数：%s 错误", idStr), Result: nil})
		return
	}

	r, err := l.listService.DeleteByID(id, userContext.(*commons.UserContext))
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}
	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: r})
}

func NewListController(service *ListService) *ListController {
	return &ListController{listService: service}
}
