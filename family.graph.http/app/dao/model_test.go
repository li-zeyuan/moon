package dao

import (
	"context"
	"testing"

	"github.com/li-zeyuan/micro/family.graph.http/app/model/inner"
	"github.com/li-zeyuan/micro/family.graph.http/library/middleware"
)

/*
updated_at设置为自动更新
ALTER TABLE member_relate
    CHANGE updated_at updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
*/

func TestCreateFamilyMember(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewFamilyMemberDao(infra.DB)

	m := familyDao.db.Migrator()
	err := m.CreateTable(&inner.FamilyMemberModel{})
	if err != nil {
		infra.Log.Error("create table error: ", err)
		return
	}

	err = m.CreateIndex(&inner.MemberRelationModel{}, "idx_deleted_at")
	if err != nil {
		infra.Log.Error("create index error: ", err)
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

func TestCreateMemberRelation(t *testing.T) {
	infra := middleware.NewInfra(context.Background(), "")
	familyDao := NewRelationDao(infra.DB)

	m := familyDao.db.Migrator()
	err := m.CreateTable(&inner.MemberRelationModel{})
	if err != nil {
		infra.Log.Error("create table error: ", err)
		return
	}

	err = m.CreateIndex(&inner.MemberRelationModel{}, "idx_uid")
	if err != nil {
		infra.Log.Error("create index error: ", err)
		return
	}
}
