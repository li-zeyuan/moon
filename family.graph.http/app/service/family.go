package service

import (
	"github.com/li-zeyuan/micro/family.graph.http/app/dao"
	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
)

var Family = familyService{}

type familyService struct{}

func (*familyService) Create(infra *middleware.Infra, req *model.FamilyAPICreateReq) error {
	familyDao := dao.NewFamilyDao(infra.DB)

	family := new(inner.FamilyModel)
	family.ID = sequence.NewID()
	family.Uid = req.Uid
	family.Name = req.Name
	family.Portrait = req.Portrait
	family.Description = req.Description
	family.Option = inner.OptionCreate
	return familyDao.Save(infra, []*inner.FamilyModel{family})
}
