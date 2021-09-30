package inner

import (
	"github.com/li-zeyuan/micro/moon.common.api/model"
)

const (
	TableNameFamily = "family"

	ColumnFamilyUid         = "uid"
	ColumnFamilyName        = "name"
	ColumnFamilyPortrait    = "portrait"
	ColumnFamilyDescription = "description"
	ColumnFamilyOption      = "option"
)

type FamilyModel struct {
	model.BaseModel
	Name        string // 家族名
	Portrait    string // 家族头像
	Description string // 描述
}
