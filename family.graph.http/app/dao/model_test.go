package dao

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

/*
updated_at设置为自动更新
ALTER TABLE member_relate
    CHANGE updated_at updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
*/

func crateIndex(db *gorm.DB, prt interface{}, index []string) error {
	if reflect.ValueOf(prt).Kind() != reflect.Ptr {
		return errors.New("prt must pointer")
	}

	m := db.Migrator()
	err := m.CreateTable(prt)
	if err != nil {
		return err
	}

	for _, i := range index {
		err = m.CreateIndex(prt, i)
		if err != nil {
			return fmt.Errorf("crate index: %s error: %s", i, err.Error())
		}
	}

	return nil
}

func TestCreateFamilyMember(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewMemberDao(infra.DB)

	m := familyDao.db.Migrator()
	err := m.CreateTable(&inner.FamilyMemberModel{})
	if err != nil {
		infra.Log.Error("create table error: ", err)
		return
	}

	err = m.CreateIndex(&inner.FamilyMemberModel{}, "idx_uid")
	if err != nil {
		infra.Log.Error("create index error: ", err)
		return
	}
	err = m.CreateIndex(&inner.FamilyMemberModel{}, "idx_family_id")
	if err != nil {
		infra.Log.Error("create index error: ", err)
		return
	}
}

func TestCreateFamily(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewFamilyDao(infra.DB)

	m := familyDao.db.Migrator()
	err := m.CreateTable(&inner.FamilyModel{})
	if err != nil {
		infra.Log.Error("create table error: ", err)
		return
	}

	err = m.CreateIndex(&inner.FamilyModel{}, "idx_deleted_at")
	if err != nil {
		infra.Log.Error("create index error: ", err)
		return
	}

	err = m.CreateIndex(&inner.FamilyModel{}, "idx_uid")
	if err != nil {
		infra.Log.Error("create index error: ", err)
		return
	}
}

func TestMigrateFamilyGraph(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewGraphDao(infra.DB)
	err := crateIndex(familyDao.db, &inner.FamilyGraphModel{}, []string{})
	assert.Equal(t, err, nil)
	t.Log(err)
}
