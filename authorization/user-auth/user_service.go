package userauth

import (
	"errors"
	"fmt"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"github.com/jiangliangquan1/iot-streaming/repository"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"github.com/sirupsen/logrus"
	"regexp"
	"time"
)

type UserService struct {
	rep        *repository.UserRepository
	logger     *logrus.Logger
	jwtManager *JwtManager
}

func (u *UserService) SignUp(req *UserSignUpRequest) (*UserSignUpResult, error) {

	err := u.validSignUpParam(req)
	if err != nil {
		u.logger.Errorf("sign-up valid param failed: param:[%v], err:[%s]", *req, err.Error())
		return nil, err
	}

	m := &models.User{
		Name:     req.Name,
		Nick:     req.Nick,
		Password: req.Password,
		Enable:   true,
		Email:    req.Email,
	}
	err = u.rep.Add(m)
	if err != nil {
		return nil, err
	}

	return &UserSignUpResult{ID: m.ID, Name: m.Name, Nick: m.Nick, Email: m.Email}, err
}

func (u *UserService) Login(req *UserLoginRequest) (*UserLoginResponse, error) {
	user, _ := u.rep.GetByName(req.UserName)
	if user == nil {
		return nil, errors.New("no such user")
	}

	if user.Password != req.Password {
		return nil, errors.New("wrong password")
	}

	tokenInfo, err := u.jwtManager.GenerateToken(user.ID, user.Name, true)
	if err != nil {
		return nil, err
	}

	return &UserLoginResponse{
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
		ExpiredIn:    tokenInfo.ExpiredIn}, nil
}

func (u *UserService) refreshToken(req *UserRefreshTokenRequest) (*UserLoginResponse, error) {

	c, err := u.jwtManager.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if c.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("refresh token expired")
	}

	tokenInfo, err := u.jwtManager.GenerateToken(c.ID, c.Username, true)
	if err != nil {
		return nil, err
	}

	return &UserLoginResponse{
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
		ExpiredIn:    tokenInfo.ExpiredIn}, nil

}

func (u *UserService) validSignUpParam(req *UserSignUpRequest) error {
	if len(req.Name) <= 0 {
		return errors.New("invalid user name")
	}

	exit, _ := u.rep.GetByName(req.Name)
	if exit != nil {
		return errors.New(fmt.Sprintf("user name: %s is already in use", req.Name))
	}

	if !validEmail(req.Email) {
		return errors.New("invalid email address")
	}

	exit, _ = u.rep.GetByEmail(req.Email)
	if exit != nil {
		return errors.New(fmt.Sprintf("user email: %s is already in use", req.Email))
	}

	//todo 完成密码校验
	//if !validPassword(req.Password, req.Name) {
	//	return errors.New("invalid password")
	//}

	return nil

}

func NewUserService(rep *repository.UserRepository, logger *logrus.Logger, jwtManager *JwtManager) *UserService {
	return &UserService{
		rep:        rep,
		logger:     logger,
		jwtManager: jwtManager,
	}
}

func validEmail(email string) bool {
	// 正则表达式用于校验邮箱格式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func validPassword(password string, userName string) bool {
	_, err := commons.AESDecrypt(password, "jlqiot-streaming")

	return err == nil
}
