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

type LoginApiSingUpReq struct {
	Passport  string `json:"passport"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Name      string `json:"name"`
}
