package inner

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableNameMock = "mock"
)

type MockModel struct {
	gorm.Model
	Field1 int64
	Field2 int64
	Field3 int64
	Field4 int64
	Field5 int64
	Time   *time.Time
}
