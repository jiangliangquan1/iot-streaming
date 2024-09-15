package userauth

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"net/http"
)

type UserController struct {
	userService *UserService
}

func (u *UserController) RegisterRoute(group *gin.RouterGroup) {
	g := group.Group("v1/users")

	g.POST("sign-up", u.signUp)
	g.POST("login", u.login)
	g.POST("refresh-token", u.refreshToken)
}

// 用户注册
func (u *UserController) signUp(ctx *gin.Context) {
	var req UserSignUpRequest
	ctx.BindJSON(&req)

	result, err := u.userService.SignUp(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}

	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: result})
}

func (u *UserController) login(ctx *gin.Context) {
	var req UserLoginRequest
	ctx.BindJSON(&req)

	result, err := u.userService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}

	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: result})
}

func (u *UserController) refreshToken(ctx *gin.Context) {
	var req UserRefreshTokenRequest
	ctx.BindJSON(&req)

	result, err := u.userService.refreshToken(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 500, Msg: err.Error(), Result: nil})
		return
	}

	ctx.JSON(http.StatusOK, commons.BaseResponse{Code: 200, Msg: "success", Result: result})

}

func NewUserController(service *UserService) *UserController {
	return &UserController{userService: service}
}
