package zlmediaserver

type CreateRequest struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	ApiBaseUrl string `json:"api_base_url"`
	ApiSecret  string `json:"api_secret"`
}

type UpdateRequest struct {
	CreateRequest
	ID int64 `json:"id"`
}

type Response struct {
	ID             int64  `json:"id"`
	ServerID       string `json:"server_id"`
	ServerName     string `json:"server_name"`
	ApiBaseUrl     string `json:"api_base_url"`
	ApiSecret      string `json:"api_secret"`
	RunningConfigs string `json:"running_configs"`
}
