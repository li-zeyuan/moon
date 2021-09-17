package inner

import (
	"time"

	basemodel "github.com/li-zeyuan/micro/micro.common.api/model"
)

const (
	TableFamilyGraph = "family_graph"

	ColumnGraphUid         = "uid"
	ColumnGraphFamilyID    = "family_id"
	ColumnGraphFatherUid   = "father_node"
	ColumnGraphSpouseUid   = "spouse_node"
	ColumnGraphIndex       = "index_num"
	ColumnGraphName        = "name"
	ColumnGraphGender      = "gender"
	ColumnGraphBirth       = "birth"
	ColumnGraphDeathTime   = "death_time"
	ColumnGraphPortrait    = "portrait"
	ColumnGraphHometown    = "hometown"
	ColumnGraphDescription = "description"
)

// FamilyGraphModel 族谱图
type FamilyGraphModel struct {
	basemodel.BaseModel
	FamilyId    int64     `gorm:"index:idx_family_id"`   // 家族id
	FatherNode  int64     `gorm:"index:idx_father_node"` // 父节点
	SpouseNode  int64     // 配偶节点
	IsSpouse    bool      // 是否配偶节点
	IndexNum    int       `gorm:"default:1"` // 兄弟节点间的排序，从1开始
	Name        string    // 姓名
	Gender      int       // 性别
	Birth       time.Time // 生日
	DeathTime   time.Time // 死亡日期
	Portrait    string    // 头像
	Hometown    string    // 家乡
	Description string    // 简介
}

type IndexObj struct {
	IndexNum int
}
