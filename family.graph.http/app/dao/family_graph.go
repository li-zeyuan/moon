package dao

import (
	"fmt"

	basemodel "github.com/li-zeyuan/micro/micro.common.api/model"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"gorm.io/gorm"
)

type GraphDao struct {
	db *gorm.DB
}

func NewGraphDao(db *gorm.DB) *GraphDao {
	return &GraphDao{
		db: db,
	}
}

func (d *GraphDao) IsExistBaseNode(infra *middleware.Infra) (bool, error) {
	var amount int64
	err := d.db.Table(inner.TableFamilyGraph).
		Count(&amount).Error
	if err != nil {
		infra.Log.Error("check if exist base note error: ", err)
		return false, err
	}

	return amount == 0, nil
}

func (d *GraphDao) Save(infra *middleware.Infra, models []*inner.FamilyGraphModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.TableFamilyGraph).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create member relation error: ", err)
		return err
	}

	return nil
}

func (d *GraphDao) GetIndex(infra *middleware.Infra, fatherUid int64) (int, error) {
	if fatherUid == 0 {
		return 0, nil
	}

	m := new(inner.IndexObj)
	err := d.db.Table(inner.TableFamilyGraph).
		Select(inner.ColumnGraphIndex).
		Where(fmt.Sprintf("%s=?", inner.ColumnGraphFatherUid), fatherUid).
		//Order(fmt.Sprintf("%s desc", inner.ColumnIndex)).
		Last(m).Error
	if err != nil {
		infra.Log.Error("get member relation index error: ", err)
		return 0, err
	}

	return m.Index, nil
}

func (d *GraphDao) NodeByIds(infra *middleware.Infra, id int64) (*inner.FamilyGraphModel, error) {
	m := new(inner.FamilyGraphModel)
	err := d.db.Table(inner.TableFamilyGraph).
		Where(fmt.Sprintf("%s = ?", basemodel.ColumnId), id).
		Where("deleted_at is null").
		Find(&m).Error
	if err != nil {
		infra.Log.Error("get node by id error: ", err)
		return nil, err
	}

	return m, nil
}

func (d *GraphDao) NodeByFamilyId(infra *middleware.Infra, familyId int64) ([]*inner.FamilyGraphModel, error) {
	models := make([]*inner.FamilyGraphModel, 0)
	err := d.db.Table(inner.TableFamilyGraph).
		Where(fmt.Sprintf("%s = ?", inner.ColumnGraphFamilyID), familyId).
		Where("deleted_at is null").
		Find(&models).Error
	if err != nil {
		infra.Log.Error("get node by family id error: ", err)
		return nil, err
	}

	return models, nil
}
