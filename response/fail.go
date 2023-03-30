package response

type FailResponse struct {
	Status    string `json:"status"`
	Message   string `json:"msg"`
	ErrorCode int    `json:"errcode"`
}
