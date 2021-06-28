package service

import (
	"context"
	"log"

	"github.com/li-zeyuan/micro/micro.common.api/errorenum"
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/user.db.rpc/app/dao"
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
		pfModel.Name = pf.Name
		pfModel.Passport = pf.Passport
		pfModel.Password = pf.Password
		profiles = append(profiles, pfModel)
	}

	userDao := dao.NewUser(infra.DB)
	err := userDao.Save(infra, profiles)
	if err != nil {
		return &profile.SaveResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	return &profile.SaveResp{}, nil
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

func (s *ProfileServer) Upsert(ctx context.Context, in *profile.UpdateReq) (*profile.UpdateResp, error) {
	log.Printf("Received: %v", in.GetProfiles())
	return &profile.UpdateResp{}, nil
}
