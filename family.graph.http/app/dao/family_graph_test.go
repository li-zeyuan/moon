package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRelationDao_Save(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewGraphDao(infra.DB)

	fModel := new(inner.FamilyGraphModel)
	err := familyDao.Save(infra, []*inner.FamilyGraphModel{fModel})
	assert.Equal(t, err, nil)
}

func TestRelationDao_GetIndex(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewGraphDao(infra.DB)

	index, err := familyDao.GetIndex(infra, 321518596442025984)
	assert.Equal(t, err, nil)
	t.Log(index)
}

func TestRelationDao_GraphRootNode(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewGraphDao(infra.DB)

	root, err := familyDao.GraphRootNode(infra, 320124380004740864)
	assert.Equal(t, err, nil)
	t.Log(root)
}

func TestRelationDao_NodeByIds(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewGraphDao(infra.DB)

	fModel := new(inner.FamilyGraphModel)
	fModel.FatherNode = 11
	index, err := familyDao.NodeByIds(infra, 1)
	assert.Equal(t, err, nil)
	t.Log(index)
}

func TestRelationDao_NodeByFamilyId(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	graphDao := NewGraphDao(infra.DB)

	fModel := new(inner.FamilyGraphModel)
	fModel.FatherNode = 11
	index, err := graphDao.NodeByFamilyId(infra, 33)
	assert.Equal(t, err, nil)
	t.Log(index)
}

func TestRelationDao_UpdateParentNode(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	graphDao := NewGraphDao(infra.DB)

	updateColumnMap := make(map[string]interface{})
	updateColumnMap[inner.ColumnGraphSpouseUid] = 22
	err := graphDao.UpdateByCurrentNode(infra, 321518596442025984, updateColumnMap)
	assert.Equal(t, err, nil)
}
