package models

type User struct {
	Common
	Name     string `json:"name" description:"用户名"`
	Nick     string `json:"nick" description:"昵称"`
	Password string `json:"password" description:"密码"`
	Enable   bool   `json:"enable" description:"启用状态"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
