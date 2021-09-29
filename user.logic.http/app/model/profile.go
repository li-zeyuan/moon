package model

type ProfileApiDetailResp struct {
	UpdatedAt int64  `json:"updated_at"`
	Uid       int64  `json:"uid"`      // 用户ID
	Name      string `json:"name"`     // 用户昵称
	Passport  string `json:"passport"` // 用户账号
	Gender    int    `json:"gender"`   // 性别 1-男；2-女
	Birth     int64  `json:"birth"`    // 生日
	Portrait  string `json:"portrait"` // 头像
	Hometown  string `json:"hometown"` // 家乡
	Phone     string `json:"phone"`    // 手机
}

type ProfileApiUpdateRep struct {
	Name     string `json:"name" validate:"max=5,min=2"` // 必填 用户昵称
	Gender   int    `json:"gender" validate:"oneof=1 2"` // 必填 性别
	Birth    int64  `json:"birth" validate:"required"`   // 必填 生日
	Portrait string `json:"portrait"`                    // 头像
	Hometown string `json:"hometown"`                    // 家乡
	Phone    string `json:"phone"`                       // 手机号
}

type UserProfileUpdateField struct {
	Uid        int64
	Name       *string // 用户昵称
	Phone      *string // 手机号
	Password   *string // 用户密码
	Gender     *int    // 性别
	Birth      *int64  // 生日
	Portrait   *string // 头像
	Hometown   *string // 家乡
	Openid     *string // WX用户openid
	SessionKey *string // session_key
}
