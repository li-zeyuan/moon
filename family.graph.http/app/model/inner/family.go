package inner

import (
	"github.com/li-zeyuan/micro/micro.common.api/model"
)

const (
	TableNameFamily = "family"

	OptionCreate = 1 // 创建
	OptionJoin   = 2 // 加入

	ColumnFamilyUid         = "uid"
	ColumnFamilyName        = "name"
	ColumnFamilyPortrait    = "portrait"
	ColumnFamilyDescription = "description"
	ColumnFamilyOption      = "option"
)

type FamilyModel struct {
	model.BaseModel
	Uid         int64  `gorm:"index:idx_uid"` // 用户ID
	Name        string // 家族名
	Portrait    string // 家族头像
	Description string // 描述
	Option      int    // 操作
}
