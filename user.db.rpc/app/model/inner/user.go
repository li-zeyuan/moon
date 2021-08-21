package inner

import (
	"time"

	"github.com/li-zeyuan/micro/micro.common.api/utils"

	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"

	"gorm.io/gorm"
)

const (
	TableNameUserProfile = "user_profile"

	ColumnUid         = "uid"
	ColumnName        = "name"
	ColumnPassport    = "passport"
	ColumnPassword    = "password"
	ColumnGender      = "gender"
	ColumnBirth       = "birth"
	ColumnPortrait    = "portrait"
	ColumnHometown    = "hometown"
	ColumnDescription = "description"
)

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

func ProfileModel2Pf(userPfs []*UserProfileModel) []*profile.Profile {
	pfs := make([]*profile.Profile, 0)
	for _, userPf := range userPfs {
		pf := new(profile.Profile)
		pf.Uid = userPf.Uid
		pf.Name = userPf.Name
		pf.Passport = userPf.Passport
		pf.Password = userPf.Password
		pf.Gender = int32(userPf.Gender)
		pf.Birth = utils.Time2TimeStamp(userPf.Birth)
		pf.Portrait = userPf.Portrait
		pf.Hometown = userPf.Hometown
		pf.Description = userPf.Description
		pfs = append(pfs, pf)
	}

	return pfs
}
