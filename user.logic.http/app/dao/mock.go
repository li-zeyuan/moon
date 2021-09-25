package dao

import (
	"gorm.io/gorm"
	"log"

	"github.com/li-zeyuan/micro/user.logic.http/app/model/inner"
)

type MockDao struct {
	db *gorm.DB
}

func NewMock(db *gorm.DB) *MockDao {
	return &MockDao{
		db: db,
	}
}

func (d *MockDao) Save(models []*inner.MockModel) error {
	err := d.db.Table(inner.TableNameMock).Create(models).Error
	if err != nil {
		log.Println("batch save error: ", err)
		return err
	}
	return nil
}
