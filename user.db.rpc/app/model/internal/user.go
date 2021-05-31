package internal

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Id       int64  // 用户ID
	Passport string // 用户账号
	Password string // 用户密码
	Nickname string // 用户昵称
}
