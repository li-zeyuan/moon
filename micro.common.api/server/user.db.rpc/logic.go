package userdbrpc

import (
	"errors"

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

func InsertProfile(baseInfra *middleware.BaseInfra, pfs []*profile.Profile) error {
	if len(pfs) == 0 {
		return nil
	}

	req := profile.SaveReq{}
	req.Profiles = pfs
	profileRpcResp := profile.SaveResp{}
	err := utils.Invoke(baseInfra, userDbRpcAddress, UrlProfileSave, &req, &profileRpcResp)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProfile(baseInfra *middleware.BaseInfra, pfs []*profile.Profile) error {
	if len(pfs) == 0 {
		return nil
	}

	req := profile.UpdateReq{}
	req.Profiles = pfs
	profileRpcResp := profile.UpdateResp{}
	err := utils.Invoke(baseInfra, userDbRpcAddress, UrlProfileUpdate, &req, &profileRpcResp)
	if err != nil {
		return err
	}

	return nil
}

func UpsertProfile(baseInfra *middleware.BaseInfra, pfUpdateField *ProfileUpdateField) error {
	if pfUpdateField.Passport == nil || len(*pfUpdateField.Passport) == 0 {
		return errors.New("passport required")
	}

	pfMap, err := GetProfileByPassport(baseInfra, []string{*pfUpdateField.Passport})
	if err != nil {
		return err
	}

	existPf, ok := pfMap[*pfUpdateField.Passport]
	if !ok {
		pf := new(profile.Profile)
		pf.Passport = *pfUpdateField.Passport
		generateProfile(pf, pfUpdateField)

		err = InsertProfile(baseInfra, []*profile.Profile{pf})
		if err != nil {
			return err
		}
	} else {
		generateProfile(existPf, pfUpdateField)
		err = UpdateProfile(baseInfra, []*profile.Profile{existPf})
		if err != nil {
			return err
		}
	}

	return nil
}

func generateProfile(pf *profile.Profile, pfUpdateField *ProfileUpdateField) {
	if pfUpdateField.Name != nil {
		pf.Name = *pfUpdateField.Name
	}
	if pfUpdateField.Password != nil {
		pf.Password = *pfUpdateField.Password
	}
	if pfUpdateField.Gender != nil {
		pf.Gender = *pfUpdateField.Gender
	}
	if pfUpdateField.Birth != nil {
		pf.Birth = *pfUpdateField.Birth
	}
	if pfUpdateField.Portrait != nil {
		pf.Portrait = *pfUpdateField.Portrait
	}
	if pfUpdateField.Hometown != nil {
		pf.Hometown = *pfUpdateField.Hometown
	}
	if pfUpdateField.Description != nil {
		pf.Description = *pfUpdateField.Description
	}
}
