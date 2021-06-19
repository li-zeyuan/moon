package dao

import (
	"testing"

	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/boot"
	"github.com/stretchr/testify/assert"
)

func TestUserDao_Save(t *testing.T) {
	infra := boot.NewInfra("")
	models := []*inner.UserModel{
		&inner.UserModel{
			Uid:       11,
			Passport: "lizeyuan",
			Password: "lizeyuan",
			Nickname: "nick",
		},
	}

	userDao := NewUser(infra.DB)
	err := userDao.Save(models)
	assert.Equal(t, err, nil)
}
