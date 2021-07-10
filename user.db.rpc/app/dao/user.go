package dao

import (
	"fmt"
	"github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/interceptor"
	"gorm.io/gorm"
	"reflect"
)

type UserDao struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (d *UserDao) Save(infra *interceptor.Infra, models []*inner.UserProfileModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.UserModelTableName).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create users error: ", err)
		return err
	}

	return nil
}

func (d *UserDao) Get(infra *interceptor.Infra, uids []int64) ([]*inner.UserProfileModel, error) {
	if len(uids) == 0 {
		return nil, nil
	}

	models := make([]*inner.UserProfileModel, 0)
	err := d.db.Table(inner.UserModelTableName).
		Where("").
		Find(&models).Error
	if err != nil {
		infra.Log.Error("get user by uids error: ", err)
		return nil, err
	}

	return models, nil
}

func (d *UserDao) GetByPassport(infra *interceptor.Infra, passports []string) ([]*inner.UserProfileModel, error) {
	if len(passports) == 0 {
		return nil, nil
	}

	models := make([]*inner.UserProfileModel, 0)
	err := d.db.Table(inner.UserModelTableName).
		Where("passport in (?)", passports).
		Find(&models).Error
	if err != nil {
		infra.Log.Error("get user by passport error: ", err)
		return nil, err
	}

	return models, nil
}

func (d *UserDao) GetColumnTypes(i interface{}) []string {
	//d.db.AutoMigrate(&inner.UserProfileModel{})

	rt := reflect.TypeOf(i).FieldByNameFunc()

	result, _ := d.db.Debug().Migrator().ColumnTypes(&inner.UserProfileModel{Uid: 1})
	for _, v := range result {
		fmt.Println(v.Name())
	}
	return nil
}
