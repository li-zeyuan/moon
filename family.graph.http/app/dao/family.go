package dao

import (
	"fmt"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	basemodel "github.com/li-zeyuan/micro/micro.common.api/model"
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

func (d *FamilyDao) OneById(infra *middleware.Infra, id int64) (*inner.FamilyModel, error) {
	m := new(inner.FamilyModel)
	err := d.db.Table(inner.TableNameFamily).
		Where(fmt.Sprintf("%s = ?", basemodel.ColumnId), id).
		Where(fmt.Sprintf("%s is null", basemodel.ColumnDeleteAt)).
		First(m).Error
	if err != nil {
		infra.Log.Error("get family by id error: ", err)
		return nil, err
	}

	return m, nil
}
