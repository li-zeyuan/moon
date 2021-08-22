package dao

import (
	"fmt"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"gorm.io/gorm"
)

type RelationDao struct {
	db *gorm.DB
}

func NewRelationDao(db *gorm.DB) *RelationDao {
	return &RelationDao{
		db: db,
	}
}

func (d *RelationDao) IsExistBaseNode(infra *middleware.Infra) (bool, error) {
	var amount int64
	err := d.db.Table(inner.TableNameMemberRelate).
		Count(&amount).Error
	if err != nil {
		infra.Log.Error("check if exist base note error: ", err)
		return false, err
	}

	return amount == 0, nil
}

func (d *RelationDao) Save(infra *middleware.Infra, models []*inner.MemberRelationModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.TableNameMemberRelate).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create member relation error: ", err)
		return err
	}

	return nil
}

func (d *RelationDao) GetIndex(infra *middleware.Infra, fatherUid int64) (int, error) {
	if fatherUid == 0 {
		return 0, nil
	}

	m := new(inner.IndexObj)
	err := d.db.Table(inner.TableNameMemberRelate).
		Select(inner.ColumnIndex).
		Where(fmt.Sprintf("%s=?", inner.ColumnFatherUid), fatherUid).
		//Order(fmt.Sprintf("%s desc", inner.ColumnIndex)).
		Last(m).Error
	if err != nil {
		infra.Log.Error("get member relation index error: ", err)
		return 0, err
	}

	return m.Index, nil
}
