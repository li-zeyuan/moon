package model

const (
	OptionAddBaseNode   = 1
	OptionAddFatherNode = 2
	OptionAddChildNode  = 3
	OptionAddSpouseNode = 4
)

type FamilyGraphAPICreateReq struct {
	Option      int    `json:"option" validate:"oneof=1 2 3 4"` // 1-添加跟节点；2-添加父节点；3-添加孩子节点；4添加配偶节点
	FamilyId    int64  `json:"family_id" validate:"gt=0"`
	CurrentNode int64  `json:"current_node"`
	FatherNode  int64  `json:"father_node"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Birth       int64  `json:"birth"`
	DeathTime   int64  `json:"death_time"`
	Portrait    string `json:"portrait"`
	Hometown    string `json:"hometown"`
	Description string `json:"description"`
}

type FamilyGraphAPIDetailReq struct {
	Node int64 `json:"node" validate:"gt=0"`
}

type FamilyGraphAPIDetailResp struct {
	Node        int64  `json:"node"`
	Name        string `json:"name"`
	IndexNum    int    `json:"index_num"`
	Gender      int    `json:"gender"`
	Birth       int64  `json:"birth"`
	DeathTime   int64  `json:"death_time"`
	Portrait    string `json:"portrait"`
	Hometown    string `json:"hometown"`
	Description string `json:"description"`
}

type FamilyGraphAPIUpdateReq struct {
	Node        int64   `json:"node" validate:"gt=0"`
	Name        *string `json:"name"`
	Gender      *int    `json:"gender"`
	Birth       *int64  `json:"birth"`
	DeathTime   *int64  `json:"death_time"`
	Portrait    *string `json:"portrait"`
	Hometown    *string `json:"hometown"`
	Description *string `json:"description"`
}

type FamilyGraphAPIDelReq struct {
	Node int64 `json:"node" validate:"gt=0"`
}

type FamilyGraphAPIGraphReq struct {
	FamilyId int64 `json:"family_id" validate:"gt=0"`
}

type FamilyGraphNode struct {
	Node        int64  `json:"node"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Birth       int64  `json:"birth"`
	DeathTime   int64  `json:"death_time"`
	Portrait    string `json:"portrait"`
	Hometown    string `json:"hometown"`
	Description string `json:"description"`
	SpouseNode  int64  `json:"spouse_node"`
}

type FamilyGraphTree struct {
	FamilyGraphNode
	Spouse   []*FamilyGraphNode `json:"spouse"`
	Children []*FamilyGraphTree `json:"children"`
}

type FamilyGraphAPIGraphResp struct {
	FamilyId int64            `json:"family_id"`
	Graph    *FamilyGraphTree `json:"graph"`
}
