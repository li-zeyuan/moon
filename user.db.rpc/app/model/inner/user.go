package inner

import (
	"gorm.io/gorm"
)

const UserModelTableName = "user_info"

type UserModel struct {
	gorm.Model
	Uid      int64  // 用户ID
	Nickname string // 用户昵称
	Passport string // 用户账号
	Password string // 用户密码
}
