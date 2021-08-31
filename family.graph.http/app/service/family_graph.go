package service

import (
	"time"

	"github.com/li-zeyuan/micro/family.graph.http/app/dao"
	"github.com/li-zeyuan/micro/family.graph.http/app/model"
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/li-zeyuan/micro/micro.common.api/errorenum"
	basemodel "github.com/li-zeyuan/micro/micro.common.api/model"
	"github.com/li-zeyuan/micro/micro.common.api/sequence"
	"github.com/li-zeyuan/micro/micro.common.api/utils"
)

var FamilyGraph = familyGraphService{}

type familyGraphService struct{}

func (g *familyGraphService) verifyCreateNode(infra *middleware.Infra, req *model.FamilyGraphAPICreateReq) error {
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
		if req.CurrentNode == 0 {
			return errorenum.ErrorCurrentParamsNode
		}
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
	switch req.Option {
	case model.OptionAddBaseNode:
		err = saveNodeByOpt(infra, sequence.NewID(), req)
		if err != nil {
			return err
		}
	case model.OptionAddFatherNode:
		fatherNode := req.FatherNode
		if req.FatherNode == 0 {
			fatherNode = sequence.NewID()
		}

		err = saveNodeByOpt(infra, fatherNode, req)
		if err != nil {
			return err
		}

		updateColumnMap := make(map[string]interface{})
		updateColumnMap[inner.ColumnGraphFatherUid] = fatherNode
		err = graphDao.UpdateByCurrentNode(infra, req.CurrentNode, updateColumnMap)
		if err != nil {
			return err
		}
	case model.OptionAddChildNode:
		err = saveNodeByOpt(infra, sequence.NewID(), req)
		if err != nil {
			return err
		}
	case model.OptionAddSpouseNode:
		spouseNode := sequence.NewID()
		err = saveNodeByOpt(infra, spouseNode, req)
		if err != nil {
			return err
		}

		updateColumnMap := make(map[string]interface{})
		updateColumnMap[inner.ColumnGraphSpouseUid] = spouseNode
		err = graphDao.UpdateByCurrentNode(infra, req.CurrentNode, updateColumnMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveNodeByOpt(infra *middleware.Infra, nodeId int64, req *model.FamilyGraphAPICreateReq) error {
	graphDao := dao.NewGraphDao(infra.DB)

	node := new(inner.FamilyGraphModel)
	node.ID = nodeId
	node.FamilyId = req.FamilyId
	node.Name = req.Name
	node.Gender = req.Gender
	node.Birth = utils.TimeStamp2Time(req.Birth)
	node.DeathTime = utils.TimeStamp2Time(req.DeathTime)
	node.Portrait = req.Portrait
	node.Hometown = req.Hometown
	node.Description = req.Description
	if req.Option == model.OptionAddChildNode {
		index, err := graphDao.GetIndex(infra, req.CurrentNode)
		if err != nil {
			return err
		}

		node.IndexNum = index + 1
		node.FatherNode = req.CurrentNode
	}

	return graphDao.Save(infra, []*inner.FamilyGraphModel{node})
}

func (g *familyGraphService) DetailNode(infra *middleware.Infra, req *model.FamilyGraphAPIDetailReq) (*model.FamilyGraphAPIDetailResp, error) {
	graphDao := dao.NewGraphDao(infra.DB)
	curNode, err := graphDao.NodeByIds(infra, req.Node)
	if err != nil {
		return nil, err
	}

	resp := new(model.FamilyGraphAPIDetailResp)
	resp.Node = curNode.ID
	resp.Name = curNode.Name
	resp.IndexNum = curNode.IndexNum
	resp.Gender = curNode.Gender
	resp.Birth = utils.Time2TimeStamp(curNode.Birth)
	resp.DeathTime = utils.Time2TimeStamp(curNode.DeathTime)
	resp.Portrait = curNode.Portrait
	resp.Hometown = curNode.Hometown
	resp.Description = curNode.Description

	return resp, nil
}

func (g *familyGraphService) UpdateNode(infra *middleware.Infra, req *model.FamilyGraphAPIUpdateReq) error {
	graphDao := dao.NewGraphDao(infra.DB)

	updateColumnMap := make(map[string]interface{})
	if req.Name != nil && len(*req.Name) > 0 {
		updateColumnMap[inner.ColumnGraphName] = *req.Name
	}
	if req.Gender != nil && *req.Gender > 0 {
		updateColumnMap[inner.ColumnGraphGender] = *req.Gender
	}
	if req.Birth != nil {
		updateColumnMap[inner.ColumnGraphBirth] = utils.TimeStamp2Time(*req.Birth)
	}
	if req.DeathTime != nil {
		updateColumnMap[inner.ColumnGraphDeathTime] = utils.TimeStamp2Time(*req.DeathTime)
	}
	if req.Portrait != nil {
		updateColumnMap[inner.ColumnGraphPortrait] = *req.Portrait
	}
	if req.Hometown != nil {
		updateColumnMap[inner.ColumnGraphHometown] = *req.Hometown
	}
	if req.Description != nil {
		updateColumnMap[inner.ColumnGraphDescription] = *req.Description
	}

	err := graphDao.UpdateByCurrentNode(infra, req.Node, updateColumnMap)
	if err != nil {
		return err
	}

	return nil
}

func (g *familyGraphService) DelNode(infra *middleware.Infra, req *model.FamilyGraphAPIDelReq) error {
	graphDao := dao.NewGraphDao(infra.DB)
	lastIndex, err := graphDao.GetIndex(infra, req.Node)
	if err != nil {
		return err
	}
	if lastIndex > 0 {
		return errorenum.ErrorOnlyDelChildNode
	}

	updateColumnMap := make(map[string]interface{})
	updateColumnMap[basemodel.ColumnDeleteAt] = time.Now()
	err = graphDao.UpdateByCurrentNode(infra, req.Node, updateColumnMap)
	if err != nil {
		return err
	}

	return nil
}
