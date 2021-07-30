package model

type LoginApiSingUpReq struct {
	Passport  string `json:"passport"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Name      string `json:"name"`
}
