package internal

import (
	"log"

	internalmodel "github.com/li-zeyuan/micro/user.db.rpc/app/model/inner"
	"github.com/li-zeyuan/micro/user.db.rpc/config"
	"gorm.io/gorm"
)

type MockDao struct {
	db *gorm.DB
}

func NewMock() *MockDao {
	return &MockDao{
		db: config.Db,
	}
}

func (d *MockDao) Save(models []*internalmodel.MockModel) error {
	err := d.db.Table(internalmodel.TableNameMock).Create(models).Error
	if err != nil {
		log.Println("batch save error: ", err)
		return err
	}
	return nil
}
