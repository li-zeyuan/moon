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
	ErrorInvalidArgument     = ErrorCode{Code: 1, Msg: "参数错误"}
	ErrorInvalidArgumentName = ErrorCode{Code: 2, Msg: "name参数错误"}
)

// login
var (
	ErrorPassportExist         = ErrorCode{Code: 200, Msg: "账号已存在"}
	ErrorPassportLetterOrDigit = ErrorCode{Code: 201, Msg: "账号只能是数字或字母"}
	ErrorPasswordLength        = ErrorCode{Code: 202, Msg: "两次输入密码输入不一致"}
	ErrorPasswordLetterOrDigit = ErrorCode{Code: 203, Msg: "密码只能是数字或字母"}
)

// family_graph
var (
	ErrorFatherUidEmpty          = ErrorCode{Code: 300, Msg: "father uid should not empty"}
	ErrorExistFamilyMember       = ErrorCode{Code: 301, Msg: "uid exist family member"}
	ErrorNotExistFamily          = ErrorCode{Code: 301, Msg: "not exist family"}
	ErrorCurrentParamsNode       = ErrorCode{Code: 302, Msg: "current node parameter error"}
	ErrorRepetitionCrateBaseNode = ErrorCode{Code: 303, Msg: "exist base node"}
	ErrorExistFatherNode         = ErrorCode{Code: 303, Msg: "exist father node"}
	ErrorOnlyDelChildNode        = ErrorCode{Code: 303, Msg: "delete child node only"}
)
