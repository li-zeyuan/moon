package inner

import "github.com/li-zeyuan/micro/micro.common.api/model"

const (
	TableNameFamilyMember = "family_member"

	OptionCreate = 1 // 创建
	OptionJoin   = 2 // 加入

	ColumnFamilyMemberUid      = "uid"
	ColumnFamilyMemberFamilyId = "family_id"
	ColumnFamilyMemberOption   = "option"
)

// FamilyMemberModel 家族成员：即可以看到族谱图的成员
type FamilyMemberModel struct {
	model.BaseModel
	Uid      int64 `gorm:"index:idx_uid"`       // 创建人id
	FamilyId int64 `gorm:"index:idx_family_id"` // 家族id
	Option   int   // 操作
}
