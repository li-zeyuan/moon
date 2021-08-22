package service

import (
	"github.com/li-zeyuan/micro/family.graph.http/app/dao"
	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/errorenum"
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
)

var Family = familyService{}

type familyService struct{}

func (*familyService) Create(infra *middleware.Infra, req *model.FamilyAPICreateReq) error {
	familyDao := dao.NewFamilyDao(infra.DB)
	familyMemberDao := dao.NewFamilyMemberDao(infra.DB)
	existFMember, err := familyMemberDao.OneByUid(infra, req.Uid)
	if err != nil {
		return err
	}
	if existFMember.ID != 0 {
		return errorenum.ErrorExistFamilyMember
	}

	family := new(inner.FamilyModel)
	family.ID = sequence.NewID()
	family.Name = req.Name
	family.Portrait = req.Portrait
	family.Description = req.Description
	err = familyDao.Save(infra, []*inner.FamilyModel{family})
	if err != nil {
		return err
	}

	familyMember := new(inner.FamilyMemberModel)
	familyMember.Uid = req.Uid
	familyMember.FamilyId = family.ID
	familyMember.Option = inner.OptionCreate
	err = familyMemberDao.Save(infra, []*inner.FamilyMemberModel{familyMember})
	if err != nil {
		return err
	}

	return nil
}

func (*familyService) Join(infra *middleware.Infra, req *model.FamilyAPIJoinReq) error {
	familyDao := dao.NewFamilyDao(infra.DB)
	f, err := familyDao.OneById(infra, req.FamilyId)
	if err != nil {
		return err
	}
	if f.ID == 0 {
		return errorenum.ErrorNotExistFamily
	}

	familyMemberDao := dao.NewFamilyMemberDao(infra.DB)
	familyMember := new(inner.FamilyMemberModel)
	familyMember.Uid = req.Uid
	familyMember.FamilyId = req.FamilyId
	familyMember.Option = inner.OptionJoin
	return familyMemberDao.Save(infra, []*inner.FamilyMemberModel{familyMember})
}
