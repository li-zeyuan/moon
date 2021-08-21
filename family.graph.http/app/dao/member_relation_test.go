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
	familyDao := NewRelation(infra.DB)

	fModel := new(inner.MemberRelationModel)
	fModel.Uid = 11
	err := familyDao.Save(infra, []*inner.MemberRelationModel{fModel})
	assert.Equal(t, err, nil)
}
