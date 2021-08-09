package userdbrpc

const (
	ServerNameUserDbRpc = "user.db.rpc"
	userDbRpcAddress    = ":7071"

	UrlProfileSave          = "/profile.ProfileService/Save"
	UrlProfileUpdate        = "/profile.ProfileService/Update"
	UrlProfileGetByPassport = "/profile.ProfileService/GetByPassport"
)

type ProfileUpdateField struct {
	Name        *string `json:"name,omitempty"`
	Passport    *string `json:"passport,omitempty"`
	Password    *string `json:"password,omitempty"`
	Gender      *int32  `json:"gender,omitempty"`
	Birth       *int64  `json:"birth,omitempty"`
	Portrait    *string `json:"portrait,omitempty"`
	Hometown    *string `json:"hometown,omitempty"`
	Description *string `json:"description,omitempty"`
}
