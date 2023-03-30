package response

import "encoding/json"

type SuccessResponse struct {
	Status    string          `json:"status"`
	Reason    string          `json:"reason,omitempty"`
	Message   string          `json:"msg,omitempty"`
	ErrorName string          `json:"errname,omitempty"`
	Value     json.RawMessage `json:"value"`
	ErrorCode int             `json:"errcode,omitempty"`
}
