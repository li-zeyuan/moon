package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/stretchr/testify/assert"
)

func TestFamilyDao_Save(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewFamilyDao(infra.DB)

	fModel := new(inner.FamilyModel)
	fModel.Name = "11"
	err := familyDao.Save(infra, []*inner.FamilyModel{fModel})
	assert.Equal(t, err, nil)
}

func TestFamilyDao_OneById(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewFamilyDao(infra.DB)

	m, err := familyDao.OneById(infra, 32012438004740864)
	assert.Equal(t, err, nil)
	t.Log(m)
}
