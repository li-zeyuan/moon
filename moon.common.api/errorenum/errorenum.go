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
	ErrorInvalidArgument     = ErrorCode{Code: 1, Msg: "invalid params"}
	ErrorInvalidArgumentName = ErrorCode{Code: 2, Msg: "name invalid params"}
)

// login
var (
	ErrorPassportExist         = ErrorCode{Code: 200, Msg: "passport exist"}
	ErrorPassportLetterOrDigit = ErrorCode{Code: 201, Msg: "passport should letter or digit"}
	ErrorPasswordInConformity  = ErrorCode{Code: 202, Msg: "password import in_conformity"}
	ErrorPasswordLetterOrDigit = ErrorCode{Code: 203, Msg: "password should letter or digit"}
)

// family_graph
var (
	ErrorFatherUidEmpty          = ErrorCode{Code: 300, Msg: "father uid should not empty"}
	ErrorExistFamilyMember       = ErrorCode{Code: 301, Msg: "uid exist family member"}
	ErrorNotExistFamily          = ErrorCode{Code: 302, Msg: "not exist family"}
	ErrorCurrentParamsNode       = ErrorCode{Code: 303, Msg: "current node parameter error"}
	ErrorRepetitionCrateRootNode = ErrorCode{Code: 304, Msg: "exist root node"}
	ErrorNotExistRootNode        = ErrorCode{Code: 305, Msg: "not exist root node"}
	ErrorExistFatherNode         = ErrorCode{Code: 306, Msg: "exist father node"}
	ErrorOnlyDelChildNode        = ErrorCode{Code: 307, Msg: "delete child node only"}
)
