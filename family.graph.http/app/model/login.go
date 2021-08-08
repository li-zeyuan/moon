package model

type FamilyGraphAPICreateReq struct {
	Phone       string `json:"phone" validate:"eq=11"`
	Name        string `json:"name" validate:"min=0,max=5"`
	Gender      int    `json:"gender" validate:"oneof=0 1 2"`
	Birth       int    `json:"birth"`
	Description string `json:"description"`
}

type LoginApiSingUpReq struct {
	Passport  string `json:"passport"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Name      string `json:"name"`
}
