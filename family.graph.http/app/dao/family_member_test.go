package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/stretchr/testify/assert"
)

func TestFamilyMemberDao_OneByUid(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewMemberDao(infra.DB)

	m, err := familyDao.OneByUid(infra, 1111)
	assert.Equal(t, err, nil)
	t.Log(m)
}

func TestFamilyMemberDao_Del(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyMemberDao := NewMemberDao(infra.DB)

	err := familyMemberDao.Del(infra, 1111, 320124380004740864)
	assert.Equal(t, err, nil)
}
