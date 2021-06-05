package internal

import (
	"gorm.io/gorm"

	internalmodel "github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
)

type UserDao struct {
	db *gorm.DB
}

func NewUser() *UserDao {
	return &UserDao{
		db: config.Db,
	}
}

func (d *UserDao) FindOne() (*internalmodel.UserModel, error) {

	return nil, nil
}
