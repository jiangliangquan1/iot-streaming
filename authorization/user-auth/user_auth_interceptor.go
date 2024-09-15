package userauth

import (
	"github.com/gin-gonic/gin"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"github.com/jiangliangquan1/iot-streaming/repository"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserAuthInterceptor struct {
	logger     *logrus.Logger
	jwtManager *JwtManager
	userRep    *repository.UserRepository
}

func (u *UserAuthInterceptor) PreHandle(c *gin.Context) (bool, int) {

	token := c.GetHeader("Authorization")

	cla, err := u.jwtManager.ParseToken(token)
	if err != nil {
		u.logger.Error(err.Error())
		return false, http.StatusUnauthorized
	}

	if cla.ExpiresAt < time.Now().Unix() {
		u.logger.Errorf("token expired: %s", token)
		return false, http.StatusUnauthorized
	}

	user, _ := u.userRep.GetByID(cla.ID)
	if user == nil {
		u.logger.Errorf("no such user: %d", cla.ID)
		return false, http.StatusUnauthorized
	}

	c.Set("userContext", &commons.UserContext{UserId: user.ID, UserName: user.Name})

	return true, http.StatusOK
}

func (u *UserAuthInterceptor) PostHandle(c *gin.Context) (bool, int) {
	return true, http.StatusOK
}

func NewUserAuthInterceptor(logger *logrus.Logger, jwtManager *JwtManager, userRep *repository.UserRepository) *UserAuthInterceptor {
	return &UserAuthInterceptor{
		logger:     logger,
		jwtManager: jwtManager,
		userRep:    userRep,
	}
}
