package model

import "time"

const (
	ColumnId       = "id"
	ColumnCreateAt = "created_at"
	ColumnUpdateAt = "updated_at"
	ColumnDeleteAt = "deleted_at"
)

type BaseModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index:idx_deleted_at"`
}
