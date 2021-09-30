package model

type BaseResponse struct {
	DmError  int         `json:"dm_error"`
	ErrorMsg string      `json:"error_msg"`
	Data     interface{} `json:"data"`
}
