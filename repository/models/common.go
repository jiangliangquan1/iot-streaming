package models

import "time"

type Common struct {
	ID         int64     `json:"id" gorm:"primaryKey" description:"主键"`
	CreateTime time.Time `json:"create_time" description:"创建时间"`
	UpdateTime time.Time `json:"update_time" description:"修改时间"`
}
