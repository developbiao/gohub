package models

import (
	"github.com/spf13/cast"
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;"json:"id,omitempty"`
}

// CommonTimestampsField timestamps
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;"json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"json:"updated_at,omitempty"`
}

// GetStringID get ID string identify
func (m BaseModel) GetStringID() string {
	return cast.ToString(m.ID)
}
