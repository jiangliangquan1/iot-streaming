package repository

import (
	"errors"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func (d *DeviceRepository) Add(device *models.Device) error {
	r := d.db.Create(device)
	return r.Error
}

func (d *DeviceRepository) Update(device *models.Device) error {
	r := d.db.Updates(device)
	return r.Error
}

func (d *DeviceRepository) DeleteByID(id int64) (*models.Device, error) {
	device := models.Device{}
	r := d.db.Delete(&device, id)

	return &device, r.Error
}

func (d *DeviceRepository) GetByID(id int64) (*models.Device, error) {
	var device models.Device
	r := d.db.Where("id = ?", id).First(&device)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &device, nil
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}
