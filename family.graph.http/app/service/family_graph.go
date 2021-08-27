package service

import (
	"regexp"

	"github.com/li-zeyuan/micro/family.graph.http/app/dao"
	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/config"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/errorenum"
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
	userdbrpc "github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc"
	"github.com/li-zeyuan/micro/micro.common.api/server/user.db.rpc/pb/profile"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

var FamilyGraph = familyGraphService{}

type familyGraphService struct{}

func (l *familyGraphService) VerifySingUp(infra *middleware.Infra, req *model.LoginApiSingUpReq) error {
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]{1,16}$`, req.Passport); !isLetterOrDigit {
		return errorenum.ErrorPassportLetterOrDigit
	}

	if len(req.Password) > 8 || len(req.Password) == 0 || req.Password != req.Password2 {
		return errorenum.ErrorPasswordLength
	}
	if isLetterOrDigit, _ := regexp.MatchString(`^[A-Za-z0-9]{1,16}$`, req.Password); !isLetterOrDigit {
		return errorenum.ErrorPasswordLetterOrDigit
	}

	// 判断账号是否存在
	profileMap, err := userdbrpc.GetProfileByPassport(infra.BaseInfra, []string{req.Passport})
	if err != nil {
		return err
	}
	if len(profileMap) > 0 {
		return errorenum.ErrorPassportExist
	}

	return nil
}

func (l *familyGraphService) SingUp(infra *middleware.Infra, req *model.LoginApiSingUpReq) error {
	pf := new(profile.Profile)
	pf.Uid = sequence.NewID()
	pf.Name = req.Name
	pf.Passport = req.Passport
	pf.Password = req.Password

	profileRpcReq := profile.SaveReq{}
	profileRpcReq.Profiles = append(profileRpcReq.Profiles, pf)
	profileRpcResp := profile.SaveResp{}
	err := utils.Invoke(infra.BaseInfra, config.GetServerClient(userdbrpc.ServerNameUserDbRpc).Address, userdbrpc.UrlProfileSave, &profileRpcReq, &profileRpcResp)
	if err != nil {
		return err
	}

	return nil
}

func (*familyGraphService) verifyCreateNode(infra *middleware.Infra, req *model.FamilyGraphAPICreateReq) error {
	graphDao := dao.NewGraphDao(infra.DB)

	switch req.Option {
	case model.OptionAddBaseNode:
		nodes, err := graphDao.NodeByFamilyId(infra, req.FamilyId)
		if err != nil {
			return err
		}
		if len(nodes) > 0 {
			return errorenum.ErrorRepetitionCrateBaseNode
		}
	case model.OptionAddFatherNode:
		curNode, err := graphDao.NodeByIds(infra, req.CurrentNode)
		if err != nil {
			return err
		}

		if curNode.FatherNode != 0 {
			return errorenum.ErrorExistFatherNode
		}
	case model.OptionAddChildNode:
		if len(req.Name) == 0 {
			return errorenum.ErrorInvalidArgumentName
		}
	case model.OptionAddSpouseNode:
		if req.CurrentNode == 0 {
			return errorenum.ErrorCurrentParamsNode
		}
		if len(req.Name) == 0 {
			return errorenum.ErrorInvalidArgumentName
		}
	}

	return nil
}

func (g *familyGraphService) CreateNode(infra *middleware.Infra, req *model.FamilyGraphAPICreateReq) error {
	err := g.verifyCreateNode(infra, req)
	if err != nil {
		return err
	}

	graphDao := dao.NewGraphDao(infra.DB)
	index, err := graphDao.GetIndex(infra, req.FatherNode)
	if err != nil {
		return err
	}

	relation := new(inner.FamilyGraphModel)
	relation.FatherNode = req.FatherNode
	relation.Index = index + 1
	return graphDao.Save(infra, []*inner.FamilyGraphModel{relation})
}
