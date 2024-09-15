package repository

import (
	"errors"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"gorm.io/gorm"
)

type ZlMediaServerRepository struct {
	db *gorm.DB
}

func (z *ZlMediaServerRepository) Add(server *models.ZlMediaServer) error {
	r := z.db.Create(server)
	return r.Error
}

func (z *ZlMediaServerRepository) Update(server *models.ZlMediaServer) error {
	r := z.db.Updates(server)
	return r.Error
}

func (z *ZlMediaServerRepository) DeleteByID(id int64) (*models.ZlMediaServer, error) {
	server := models.ZlMediaServer{}
	r := z.db.Delete(&server, id)

	return &server, r.Error
}

func (z *ZlMediaServerRepository) GetByID(id int64) (*models.ZlMediaServer, error) {
	var server models.ZlMediaServer
	r := z.db.Where("id = ?", id).First(&server)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &server, nil
}

func (z *ZlMediaServerRepository) GetByMediaServerID(serverID string) (*models.ZlMediaServer, error) {
	var server models.ZlMediaServer
	r := z.db.Where("server_id = ?", serverID).First(&server)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &server, nil
}

func NewZlMediaServerRepository(db *gorm.DB) *ZlMediaServerRepository {
	return &ZlMediaServerRepository{db: db}
}
