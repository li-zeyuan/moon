package userdbrpc

import (
	"github.com/li-zeyuan/micro/micro.common.api/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

func GetProfileByPassport(baseInfra *middleware.BaseInfra, passports []string) (map[string]*profile.Profile, error) {
	req := profile.GetByPassportReq{}
	req.Passports = append(req.Passports, passports...)
	resp := profile.GetByPassportResp{}
	err := utils.Invoke(baseInfra, userDbRpcAddress, UrlProfileGetByPassport, &req, &resp)
	if err != nil {
		return nil, err
	}

	passportMap := make(map[string]*profile.Profile)
	for _, p := range resp.Data.List {
		passportMap[p.Passport] = p
	}

	return passportMap, nil
}
