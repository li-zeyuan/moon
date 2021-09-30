package service

import (
	"context"
	"encoding/json"

	"github.com/li-zeyuan/micro/moon.common.api/errorenum"
	"github.com/li-zeyuan/micro/moon.common.api/sequence"
	"github.com/li-zeyuan/micro/moon.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/moon.common.api/utils"
	"github.com/li-zeyuan/micro/user.db.rpc/app/dao"
	"github.com/li-zeyuan/micro/user.db.rpc/app/model"
	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/interceptor"
)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
}

func (s *ProfileServer) Save(ctx context.Context, in *profile.SaveReq) (*profile.SaveResp, error) {
	infra := interceptor.GetInfra(ctx)

	profiles := make([]*inner.UserProfileModel, 0)
	for _, pf := range in.Profiles {
		pfModel := new(inner.UserProfileModel)
		pfModel.Uid = pf.Uid
		if pfModel.Uid == 0 {
			pfModel.Uid = sequence.NewID()
		}
		pfModel.Name = pf.Name
		pfModel.Passport = pf.Passport
		pfModel.Password = pf.Password
		pfModel.Gender = int(pf.Gender)
		pfModel.Birth = utils.TimeStamp2Time(pf.Birth)
		pfModel.Portrait = pf.Portrait
		pfModel.Hometown = pf.Hometown
		pfModel.Description = pf.Description
		profiles = append(profiles, pfModel)
	}

	userDao := dao.NewUser(infra.DB)
	err := userDao.Save(infra, profiles)
	if err != nil {
		return &profile.SaveResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	bPf, err := json.Marshal(inner.ProfileModel2Pf(profiles))
	if err != nil {
		return &profile.SaveResp{DmError: -1, ErrorMsg: err.Error()}, err
	}
	return &profile.SaveResp{Data: bPf}, nil
}

func (s *ProfileServer) Update(ctx context.Context, in *profile.UpdateReq) (*profile.UpdateResp, error) {
	infra := interceptor.GetInfra(ctx)

	userDao := dao.NewUser(infra.DB)
	err := userDao.Update(infra, model.Profiles2Model(in.Profiles))
	if err != nil {
		return &profile.UpdateResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	return &profile.UpdateResp{}, nil
}

func (s *ProfileServer) Del(ctx context.Context, in *profile.DelRep) (*profile.DelResp, error) {
	return &profile.DelResp{}, nil
}

func (s *ProfileServer) Get(ctx context.Context, in *profile.GetReq) (*profile.GetResp, error) {
	infra := interceptor.GetInfra(ctx)
	if len(in.Uids) == 0 {
		return &profile.GetResp{DmError: int32(errorenum.ErrorInvalidArgument.Code), ErrorMsg: errorenum.ErrorInvalidArgument.Msg}, nil
	}

	userDao := dao.NewUser(infra.DB)
	users, err := userDao.Get(infra, in.GetUids())
	if err != nil {
		return &profile.GetResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	data := new(profile.GetRespList)
	data.List = make([]*profile.Profile, 0)
	for _, u := range users {
		uInfo := new(profile.Profile)
		uInfo.Uid = u.Uid
		uInfo.Name = u.Name
		uInfo.Passport = u.Passport
		uInfo.Password = u.Password
		data.List = append(data.List, uInfo)
	}

	return &profile.GetResp{Data: data}, nil
}

func (s *ProfileServer) GetByPassport(ctx context.Context, in *profile.GetByPassportReq) (*profile.GetByPassportResp, error) {
	infra := interceptor.GetInfra(ctx)
	if len(in.Passports) == 0 {
		return &profile.GetByPassportResp{DmError: int32(errorenum.ErrorInvalidArgument.Code), ErrorMsg: errorenum.ErrorInvalidArgument.Msg}, nil
	}

	userDao := dao.NewUser(infra.DB)
	users, err := userDao.GetByPassport(infra, in.GetPassports())
	if err != nil {
		return &profile.GetByPassportResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	data := new(profile.GetRespList)
	data.List = make([]*profile.Profile, 0)
	for _, u := range users {
		uInfo := new(profile.Profile)
		uInfo.Uid = u.Uid
		uInfo.Name = u.Name
		uInfo.Passport = u.Passport
		uInfo.Password = u.Password
		data.List = append(data.List, uInfo)
	}

	return &profile.GetByPassportResp{Data: data}, nil
}
