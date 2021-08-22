package model

type FamilyAPICreateReq struct {
	Uid         int64  `json:"uid" validate:"gt=0"`
	Name        string `json:"name" validate:"min=1,max=10"`
	Portrait    string `json:"portrait"`
	Description string `json:"description"`
}
