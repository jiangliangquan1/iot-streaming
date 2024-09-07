package commons

type BaseResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result any    `json:"result"`
}
