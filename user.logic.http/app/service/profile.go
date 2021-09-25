package service

import (
	"github.com/li-zeyuan/micro/micro.common.api/utils"
	"github.com/li-zeyuan/micro/user.logic.http/app/dao"
	"github.com/li-zeyuan/micro/user.logic.http/app/model"
	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
)

var Profile = profileService{}

type profileService struct{}

func (l *profileService) Detail(infra *middleware.Infra, uid int64) (*model.ProfileApiDetailResp, error) {
	userDao := dao.NewUser(infra.DB)
	userProfile, err := userDao.GetOne(infra, uid)
	if err != nil {
		return nil, err
	}

	resp := new(model.ProfileApiDetailResp)
	resp.UpdatedAt = utils.Time2TimeStamp(userProfile.UpdatedAt)
	resp.Uid = userProfile.Uid
	resp.Name = userProfile.Name
	resp.Passport = userProfile.Passport
	resp.Gender = userProfile.Gender
	resp.Birth = utils.Time2TimeStamp(userProfile.Birth)
	resp.Portrait = userProfile.Portrait
	resp.Hometown = userProfile.Hometown
	resp.Description = userProfile.Description

	return resp, nil
}
