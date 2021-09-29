package inner

import (
	"time"

	"github.com/li-zeyuan/micro/micro.common.api/model"
)

const (
	TableNameUserProfile = "user_profile"

	ColumnUid      = "uid"
	ColumnName     = "name"
	ColumnPhone    = "phone"
	ColumnPassport = "passport"
	ColumnPassword = "password"
	ColumnGender   = "gender"
	ColumnBirth    = "birth"
	ColumnPortrait = "portrait"
	ColumnHometown = "hometown"
)

const (
	GenderMan   = 1
	GenderWoman = 2
)

type UserProfileModel struct {
	model.BaseModel
	Uid        int64     // 用户ID
	Name       string    // 用户昵称
	Phone      string    // 手机号
	Passport   string    // 用户账号
	Password   string    // 用户密码
	Gender     int       // 性别
	Birth      time.Time // 生日
	Portrait   string    // 头像
	Hometown   string    // 家乡
	Openid     string    // WX用户openid
	SessionKey string    // session_key
}
