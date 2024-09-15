package models

type ZlMediaServer struct {
	Common
	ServerID       string `json:"server_id"`
	ServerName     string `json:"server_name"`
	UserID         int64  `json:"user_id"`
	RunningConfigs string `json:"running_configs"`
	ApiBaseUrl     string `json:"api_base_url"`
	ApiSecret      string `json:"api_secret"`
}

func (ZlMediaServer) TableName() string {
	return "zlmediaserver"
}
