package model

import (
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
)

func Profiles2Model(profiles []*profile.Profile) []*inner.UserProfileModel {
	if len(profiles) == 0 {
		return nil
	}

	pfModels := make([]*inner.UserProfileModel, 0)
	for _, pf := range profiles {
		pfModel := new(inner.UserProfileModel)
		pfModel.Uid = pf.Uid
		pfModel.Name = pf.Name
		pfModel.Passport = pf.Passport
		pfModel.Password = pf.Password
		pfModel.Gender = int(pf.Gender)
		pfModel.Birth = utils.TimeStamp2Time(pf.Birth)
		pfModel.Portrait = pf.Portrait
		pfModel.Hometown = pf.Hometown
		pfModel.Description = pf.Description
		pfModels = append(pfModels, pfModel)
	}

	return pfModels
}
