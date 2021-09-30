package userdbrpc

import (
	"encoding/json"
	"errors"

	"github.com/li-zeyuan/micro/moon.common.api/middleware"
	"github.com/li-zeyuan/micro/moon.common.api/server/user.db.rpc/pb/profile"
)

func GetProfileByPassport(baseInfra *middleware.BaseInfra, passports []string) (map[string]*profile.Profile, error) {
	req := profile.GetByPassportReq{}
	req.Passports = append(req.Passports, passports...)
	resp := profile.GetByPassportResp{}
	err := middleware.Invoke(baseInfra, userDbRpcAddress, UrlProfileGetByPassport, &req, &resp)
	if err != nil {
		return nil, err
	}

	passportMap := make(map[string]*profile.Profile)
	for _, p := range resp.Data.List {
		passportMap[p.Passport] = p
	}

	return passportMap, nil
}

func CreateProfile(baseInfra *middleware.BaseInfra, pfs []*profile.Profile) ([]*profile.Profile, error) {
	if len(pfs) == 0 {
		return nil, nil
	}

	req := profile.SaveReq{}
	req.Profiles = pfs
	profileRpcResp := profile.SaveResp{}
	err := middleware.Invoke(baseInfra, userDbRpcAddress, UrlProfileSave, &req, &profileRpcResp)
	if err != nil {
		return nil, err
	}

	resultPfs := make([]*profile.Profile, 0)
	err = json.Unmarshal(profileRpcResp.Data, &resultPfs)
	if err != nil {
		baseInfra.Log.Error("json unmarshal profile error: ", err)
		return nil, err
	}

	return resultPfs, nil
}

func UpdateProfile(baseInfra *middleware.BaseInfra, pfs []*profile.Profile) error {
	if len(pfs) == 0 {
		return nil
	}

	req := profile.UpdateReq{}
	req.Profiles = pfs
	profileRpcResp := profile.UpdateResp{}
	err := middleware.Invoke(baseInfra, userDbRpcAddress, UrlProfileUpdate, &req, &profileRpcResp)
	if err != nil {
		return err
	}

	return nil
}

func UpsertProfile(baseInfra *middleware.BaseInfra, pfUpdateField *ProfileUpdateField) (*profile.Profile, error) {
	if pfUpdateField.Passport == nil || len(*pfUpdateField.Passport) == 0 {
		return nil, errors.New("passport required")
	}

	pfMap, err := GetProfileByPassport(baseInfra, []string{*pfUpdateField.Passport})
	if err != nil {
		return nil, err
	}

	existPf, ok := pfMap[*pfUpdateField.Passport]
	if !ok {
		pf := new(profile.Profile)
		pf.Passport = *pfUpdateField.Passport
		fieldUpdate2Profile(pf, pfUpdateField)

		pfs, err := CreateProfile(baseInfra, []*profile.Profile{pf})
		if err != nil {
			return nil, err
		}
		if len(pfs) > 0 {
			existPf = pfs[0]
		}
	} else {
		fieldUpdate2Profile(existPf, pfUpdateField)
		err = UpdateProfile(baseInfra, []*profile.Profile{existPf})
		if err != nil {
			return nil, err
		}
	}

	return existPf, nil
}

func fieldUpdate2Profile(pf *profile.Profile, pfUpdateField *ProfileUpdateField) {
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
