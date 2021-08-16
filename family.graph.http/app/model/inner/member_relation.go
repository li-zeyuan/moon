package inner

import (
	"gorm.io/gorm"
)

const (
	TableNameMemberRelate = "member_relate"

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

type MemberRelationModel struct {
	gorm.Model
	Uid       int64 // 用户ID
	ParentUid int64 // 父节点uid
}
