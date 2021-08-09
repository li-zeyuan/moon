package inner

import (
	"time"

	"gorm.io/gorm"
)

const UserModelTableName = "user_info"

type UserProfileModel struct {
	gorm.Model
	Uid         int64     // 用户ID
	Name        string    // 用户昵称
	Passport    string    // 用户账号
	Password    string    // 用户密码
	Gender      int       // 性别
	Birth       time.Time // 生日
	Portrait    string    // 头像
	Hometown    string    // 家乡
	Description string    // 简介
}
