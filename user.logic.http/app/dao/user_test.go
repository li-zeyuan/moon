package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/user.logic.http/library/middleware"
	"github.com/stretchr/testify/assert"
)

func TestUserDao_Save(t *testing.T) {

}

func TestUserDao_Update(t *testing.T) {

}

func TestUserDao_Get(t *testing.T) {

}

func TestUserDao_GetByOpenid(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")

	userDao := NewUser(infra.DB)
	userInfos, err := userDao.GetByOpenid(infra, "uids")
	assert.Equal(t, err, nil)
	t.Log(userInfos)
}

func TestUserDao_GetOne(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")

	userDao := NewUser(infra.DB)
	userInfos, err := userDao.GetOne(infra, 318861102280601344)
	assert.Equal(t, err, nil)
	t.Log(userInfos)
}

func TestUserDao_GetByPassport(t *testing.T) {

}
