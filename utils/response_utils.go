package utils

type ResponseUtils struct {
	Message    string      `json:"message"`
	Response   interface{} `json:"response"`
	StatusCode int         `json:"status_code"`
}
