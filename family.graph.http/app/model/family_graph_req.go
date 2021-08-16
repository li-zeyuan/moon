package model

type FamilyGraphAPICreateReq struct {
	Passport    string `json:"Passport" validate:"required"`
	Name        string `json:"name" validate:"min=0,max=5"`
	Gender      int32  `json:"gender" validate:"oneof=1 2"`
	Birth       int64  `json:"birth"`
	Description string `json:"description"`
}

type LoginApiSingUpReq struct {
	Passport  string `json:"passport"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Name      string `json:"name"`
}
