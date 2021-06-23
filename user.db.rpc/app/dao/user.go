package dao

import (
	"log"

	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (d *UserDao) Save(models []*inner.UserProfileModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.UserModelTableName).Create(&models).Error
	if err != nil {
		log.Println("create users error: ", err)
		return err
	}

	return nil
}

func (d *UserDao) FindOne() (*inner.UserProfileModel, error) {

	return nil, nil
}
