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

func (l *profileService) Update(infra *middleware.Infra, uid int64, updateField *model.ProfileApiUpdateRep) error {
	userDao := dao.NewUser(infra.DB)
	_, err := userDao.GetOne(infra, uid)
	if err != nil {
		return err
	}

	updateItem := new(model.UserProfileUpdateField)
	updateItem.Uid = uid
	updateItem.Name = &updateField.Name
	updateItem.Gender = &updateField.Gender
	updateItem.Birth = &updateField.Birth
	if len(updateField.Portrait) > 0 {
		updateItem.Portrait = &updateField.Portrait
	}
	if len(updateField.Hometown) > 0 {
		updateItem.Hometown = &updateField.Hometown
	}
	if len(updateField.Description) > 0 {
		updateItem.Description = &updateField.Description
	}

	err = userDao.Update(infra, []*model.UserProfileUpdateField{updateItem})
	if err != nil {
		return err
	}

	return nil
}
