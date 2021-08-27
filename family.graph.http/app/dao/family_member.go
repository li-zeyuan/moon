package dao

import (
	"fmt"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"gorm.io/gorm"
)

type FamilyMemberDao struct {
	db *gorm.DB
}

func NewFamilyMemberDao(db *gorm.DB) *FamilyMemberDao {
	return &FamilyMemberDao{
		db: db,
	}
}

func (d *FamilyMemberDao) OneByUid(infra *middleware.Infra, uid int64) (*inner.FamilyMemberModel, error) {
	m := new(inner.FamilyMemberModel)
	err := d.db.Table(inner.TableNameFamilyMember).
		Where(fmt.Sprintf("%s = ?", inner.ColumnFamilyMemberUid), uid).
		Where("deleted_at is null").
		Find(m).Error
	if err != nil {
		infra.Log.Error("get family member by uid error: ", err)
		return nil, err
	}

	return m, nil
}

func (d *FamilyMemberDao) Save(infra *middleware.Infra, models []*inner.FamilyMemberModel) error {
	if len(models) == 0 {
		return nil
	}

	err := d.db.Table(inner.TableNameFamilyMember).
		Create(&models).Error
	if err != nil {
		infra.Log.Error("create family member error: ", err)
		return err
	}

	return nil
}

func (d *FamilyMemberDao) Del(infra *middleware.Infra, uid, familyId int64) error {
	err := d.db.Table(inner.TableNameFamilyMember).
		Where(fmt.Sprintf("%s = ?", inner.ColumnFamilyMemberUid), uid).
		Where(fmt.Sprintf("%s = ?", inner.ColumnFamilyMemberFamilyId), familyId).
		Delete(&inner.FamilyMemberModel{}).Error
	if err != nil {
		infra.Log.Error("del family member error: ", err)
		return err
	}

	return nil
}
