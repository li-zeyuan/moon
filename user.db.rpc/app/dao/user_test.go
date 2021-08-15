package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/interceptor"
	"github.com/stretchr/testify/assert"
)

func TestUserDao_Save(t *testing.T) {
	infra := interceptor.NewInfra(context.Background(), "")
	models := []*inner.UserProfileModel{
		&inner.UserProfileModel{
			Uid:      11,
			Passport: "lizeyuan",
			Password: "lizeyuan",
			Name:     "nick",
		},
	}

	userDao := NewUser(infra.DB)
	err := userDao.Save(infra, models)
	assert.Equal(t, err, nil)
}

func TestUserDao_Update(t *testing.T) {
	infra := interceptor.NewInfra(context.Background(), "")
	values := map[string]interface{}{
		"passport": "lizeyuan",
	}

	userDao := NewUser(infra.DB)
	err := userDao.Update(infra, 1, values)
	assert.Equal(t, err, nil)
}

func TestUserDao_Get(t *testing.T) {
	infra := interceptor.NewInfra(context.Background(), "")
	uids := []int64{
		111,
	}

	userDao := NewUser(infra.DB)
	userInfos, err := userDao.Get(infra, uids)
	assert.Equal(t, err, nil)
	t.Log(userInfos)
}

func TestUserDao_GetByPassport(t *testing.T) {
	infra := interceptor.NewInfra(context.Background(), "")
	passports := []string{
		"lizeyuan",
	}

	userDao := NewUser(infra.DB)
	userInfos, err := userDao.GetByPassport(infra, passports)
	assert.Equal(t, err, nil)
	t.Log(userInfos)
}
