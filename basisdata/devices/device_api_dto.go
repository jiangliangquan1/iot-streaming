package devices

type DeviceCreateRequest struct {
	Name        string `json:"name"`
	SecretKey   string `json:"secret_key"`
	Enable      bool   `json:"enable"`
	ProductKey  string `json:"product_key"`
	Tags        any    `json:"tags"`
	PID         *int64 `json:"pid"`
	Description string `json:"description"`
}

type DeviceUpdateRequest struct {
	DeviceCreateRequest
	ID int64 `json:"id"`
}

type DeviceResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	SecretKey   string `json:"secret_key"`
	Enable      bool   `json:"enable"`
	ProductKey  string `json:"product_key"`
	Tags        any    `json:"tags"`
	PID         *int64 `json:"pid"`
	Description string `json:"description"`
}
