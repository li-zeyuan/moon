package dao

import (
	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"gorm.io/gorm"
)

type RelationDao struct {
	db *gorm.DB
}

func NewRelation(db *gorm.DB) *RelationDao {
	return &RelationDao{
		db: db,
	}
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
