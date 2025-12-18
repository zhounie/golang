package models

import (
	"myapp/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID int64 `gorm:"primaryKey" json:"id,string"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type DeleteRequest struct {
	IDS []int64 `json:"ids" binding:"required"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == 0 {
		b.ID = utils.GenID()
	}
	return
}
