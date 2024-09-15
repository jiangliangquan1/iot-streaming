package zlmediaserver

import (
	"errors"
	"fmt"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"github.com/jiangliangquan1/iot-streaming/repository"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"github.com/sirupsen/logrus"
)

// zlmediaserver 列表管理服务
type ListService struct {
	rep    *repository.ZlMediaServerRepository
	logger *logrus.Logger
}

func (l *ListService) Create(req *CreateRequest, userContext *commons.UserContext) (*Response, error) {
	m := GetModel(req, userContext)

	err := l.rep.Add(m)
	if err != nil {
		return nil, err
	}

	return GetResponse(m), err
}

func (l *ListService) Update(req *UpdateRequest, userContext *commons.UserContext) (*Response, error) {

	exist, err := l.rep.GetByID(req.ID)
	if exist == nil || exist.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("zlmediaserver：%d 不存在", req.ID))
	}

	m := GetModel(&req.CreateRequest, userContext)
	m.ID = req.ID

	err = l.rep.Update(m)
	if err != nil {
		return nil, err
	}
	return GetResponse(m), err
}

func (l *ListService) GetByID(id int64, userContext *commons.UserContext) (*Response, error) {
	r, err := l.rep.GetByID(id)
	if err != nil {
		return nil, err
	}

	if r == nil || r.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("zlmediaserver：%d 不存在", id))
	}
	return GetResponse(r), nil
}

func (l *ListService) GetByServerID(serverID string, userContext *commons.UserContext) (*Response, error) {
	r, err := l.rep.GetByMediaServerID(serverID)
	if err != nil {
		return nil, err
	}

	if r == nil || r.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("zlmediaserver：%s 不存在", serverID))
	}
	return GetResponse(r), nil
}

func (l *ListService) DeleteByID(id int64, userContext *commons.UserContext) (*Response, error) {
	exist, err := l.rep.GetByID(id)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, nil
	}

	if exist.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("zlmediaserver：%d 不存在", id))
	}

	_, err = l.rep.DeleteByID(id)
	if err != nil {
		return nil, err
	}

	return GetResponse(exist), nil
}

func GetModel(req *CreateRequest, userContext *commons.UserContext) *models.ZlMediaServer {
	return &models.ZlMediaServer{
		ServerID:   req.ServerID,
		ServerName: req.ServerName,
		UserID:     userContext.UserId,
		ApiBaseUrl: req.ApiBaseUrl,
		ApiSecret:  req.ApiSecret,
	}
}

func GetResponse(model *models.ZlMediaServer) *Response {
	return &Response{
		ID:             model.ID,
		ServerID:       model.ServerID,
		ServerName:     model.ServerName,
		ApiBaseUrl:     model.ApiBaseUrl,
		ApiSecret:      model.ApiSecret,
		RunningConfigs: model.RunningConfigs,
	}
}

func NewListService(rep *repository.ZlMediaServerRepository, logger *logrus.Logger) *ListService {
	return &ListService{rep: rep, logger: logger}
}
