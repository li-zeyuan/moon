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

	fModel := new(inner.FamilyGraphModel)
	fModel.FatherNode = 11
	index, err := familyDao.GetIndex(infra, 111)
	assert.Equal(t, err, nil)
	t.Log(index)
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
