package errorenum

import "encoding/json"

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err ErrorCode) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err ErrorCode) HasError() bool {
	return err.Code != 0
}

// common
var (
	ErrorInvalidArgument = ErrorCode{Code: 1, Msg: "参数错误"}
)

// login
var (
	ErrorPassportLength        = ErrorCode{Code: 200, Msg: "账号长度超过限制"}
	ErrorPassportLetterOrDigit = ErrorCode{Code: 201, Msg: "账号只能是数字或字母"}
	ErrorPasswordLength        = ErrorCode{Code: 202, Msg: "两次输入密码输入不一致"}
	ErrorPasswordLetterOrDigit = ErrorCode{Code: 203, Msg: "密码只能是数字或字母"}
)
