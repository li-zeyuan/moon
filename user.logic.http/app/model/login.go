package model

type LoginApiPhoneLoginReq struct {
	Passport  string `json:"passport"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Name      string `json:"name"`
}

type LoginApiWeChatLoginReq struct {
	Code string `json:"code" validate:"min=1"`
}

type LoginApiWeChatLoginResp struct {
	Token string `json:"token"`
}

type WXSessionRet struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
