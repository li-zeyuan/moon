package service

import (
	"context"
	"github.com/li-zeyuan/micro/user.db.rpc/app/dao"
	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/boot"
	"github.com/li-zeyuan/micro/user.db.rpc/pb/profile"
	"log"

)

type ProfileServer struct {
	profile.UnimplementedProfileServiceServer
}

func (s *ProfileServer) Save(ctx context.Context, in *profile.SaveReq) (*profile.SaveResp, error) {
	infra := boot.GetInfra(ctx)

	profiles :=  make([]*inner.UserProfileModel, 0)
	for _, pf := range in.Profiles {
		pfModel := new(inner.UserProfileModel)
		pfModel.Uid = pf.Uid
		pfModel.Name = pf.Name
		pfModel.Passport = pf.Passport
		pfModel.Password = pf.Password
		profiles = append(profiles, pfModel)
	}

	userDao := dao.NewUser(infra.DB)
	err:= userDao.Save(profiles)
	if err != nil {
		return &profile.SaveResp{DmError: -1, ErrorMsg: err.Error()}, err
	}

	return &profile.SaveResp{}, nil
}

func (s *ProfileServer) Upsert(ctx context.Context, in *profile.UpdateReq) (*profile.UpdateResp, error) {
	log.Printf("Received: %v", in.GetProfiles())
	return &profile.UpdateResp{}, nil
}
