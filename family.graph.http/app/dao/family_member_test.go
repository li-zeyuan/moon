package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/stretchr/testify/assert"
)

func TestFamilyDao_OneByUid(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewFamilyMemberDao(infra.DB)

	// todo
	m, err := familyDao.OneByUid(infra, 1111)
	assert.Equal(t, err, nil)
	t.Log(m)
}
