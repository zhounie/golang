package models

import (
	"gorm.io/gorm"
)

type Product struct {
	Name string `gorm:"size:50;not null" json:"name"`
	Price float64 `gorm:"not null" json:"price"`
	gorm.Model
}
