package models

import "gorm.io/datatypes"

type Device struct {
	Common
	Name        string         `json:"name"`
	SecretKey   string         `json:"secret_key"`
	Enable      bool           `json:"enable"`
	ProductKey  string         `json:"product_key"`
	UserID      int64          `json:"user_id"`
	Tags        datatypes.JSON `json:"tags"`
	PID         *int64         `json:"pid" gorm:"column:pid"`
	Description string         `json:"description"`
}

func (*Device) TableName() string {
	return "device"
}
