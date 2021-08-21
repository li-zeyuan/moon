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
	Uid       int64 `gorm:"index:idx_uid"` // 用户ID
	FatherUid int64 // 父节点uid
	SpouseUid int64 // 配偶uid
	Index     int   // 兄弟节点间的排序，default 1
}
