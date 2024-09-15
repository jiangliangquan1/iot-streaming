package repository

import (
	"errors"
	"github.com/jiangliangquan1/iot-streaming/repository/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Add(user *models.User) error {
	result := u.db.Create(user)
	return result.Error
}

func (u *UserRepository) GetByID(id int64) (*models.User, error) {
	var user models.User
	r := u.db.Where("id = ?", id).First(&user)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &user, nil
}

func (u *UserRepository) GetByName(name string) (*models.User, error) {
	var user models.User
	r := u.db.Where("name = ?", name).First(&user)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &user, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	r := u.db.Where("email = ?", email).First(&user)

	if r.Error != nil && errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil //未查询到结果
	} else if r.Error != nil {
		return nil, r.Error //查询错误
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
