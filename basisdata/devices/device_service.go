package devices

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jiangliangquan1/iot-streaming/commons"
	"github.com/jiangliangquan1/iot-streaming/repository"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"github.com/sirupsen/logrus"
)

type DeviceService struct {
	deviceRep *repository.DeviceRepository
	logger    *logrus.Logger
}

func (d *DeviceService) Create(req *DeviceCreateRequest, userContext *commons.UserContext) (*DeviceResponse, error) {

	deviceModel := GetModel(req, userContext)

	if len(deviceModel.SecretKey) <= 0 {
		deviceModel.SecretKey = uuid.NewString()
	}

	err := d.deviceRep.Add(deviceModel)
	if err != nil {
		return nil, err
	}

	return GetResponse(deviceModel), err
}

func (d *DeviceService) Update(req *DeviceUpdateRequest, userContext *commons.UserContext) (*DeviceResponse, error) {

	exist, err := d.deviceRep.GetByID(req.ID)
	if exist == nil || exist.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("设备：%d 不存在", req.ID))
	}

	deviceModel := GetModel(&req.DeviceCreateRequest, userContext)
	deviceModel.ID = req.ID
	if len(deviceModel.SecretKey) <= 0 {
		deviceModel.SecretKey = uuid.NewString()
	}

	err = d.deviceRep.Update(deviceModel)
	if err != nil {
		return nil, err
	}

	return GetResponse(deviceModel), err
}

func (d *DeviceService) GetByID(id int64, userContext *commons.UserContext) (*DeviceResponse, error) {
	r, err := d.deviceRep.GetByID(id)
	if err != nil {
		return nil, err
	}

	if r == nil || r.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("设备：%d 不存在", id))
	}
	return GetResponse(r), nil
}

func (d *DeviceService) DeleteByID(id int64, userContext *commons.UserContext) (*DeviceResponse, error) {
	exist, err := d.deviceRep.GetByID(id)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, nil
	}

	if exist.UserID != userContext.UserId {
		return nil, errors.New(fmt.Sprintf("设备：%d 不存在", id))
	}

	_, err = d.deviceRep.DeleteByID(id)
	if err != nil {
		return nil, err
	}

	return GetResponse(exist), nil
}

func NewDeviceService(deviceRep *repository.DeviceRepository, logger *logrus.Logger) *DeviceService {
	return &DeviceService{deviceRep: deviceRep, logger: logger}
}

func GetModel(req *DeviceCreateRequest, userContext *commons.UserContext) *models.Device {

	m := &models.Device{
		Name:        req.Name,
		SecretKey:   req.SecretKey,
		Enable:      req.Enable,
		PID:         req.PID,
		Description: req.Description,
		ProductKey:  req.ProductKey,
		UserID:      userContext.UserId,
	}

	if req.Tags != nil {
		m.Tags, _ = json.Marshal(req.Tags)
	}

	return m
}

func GetResponse(model *models.Device) *DeviceResponse {
	return &DeviceResponse{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		SecretKey:   model.SecretKey,
		Enable:      model.Enable,
		PID:         model.PID,
		Tags:        model.Tags,
		ProductKey:  model.ProductKey,
	}
}
