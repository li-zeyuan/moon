package dao

import (
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"gorm.io/gorm"
)

type FamilyDao struct {
	db *gorm.DB
}

func NewFamilyDao(db *gorm.DB) *FamilyDao {
	return &FamilyDao{
		db: db,
	}
}

func (d *FamilyDao) Save(infra *middleware.Infra, models []*inner.FamilyModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.TableNameFamily).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create family error: ", err)
		return err
	}

	return nil
}
