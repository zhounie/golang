package models

import (
	"time"
	"gorm.io/gorm"
)

const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusBanned   = "banned"
)

const (
	UserRoleCustomer = "customer"
	UserRoleAdmin    = "admin"
)

type User struct {
	Username string `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	Phone    string `gorm:"size:20" json:"phone"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Role     string `gorm:"size:20;default:'customer'" json:"role"`
	Status   string `gorm:"size:20;default:'active'" json:"status"`

	Nickname  string     `gorm:"size:100" json:"nickname"`
	Gender    string     `gorm:"size:10" json:"gender"`
	Birthdate *time.Time `json:"birthdate"`

	LastLoginAt *time.Time `json:last_login_at`

	// Products []Product `gorm:"foreignKey:UserID" json:"pruducts,omitempty"`
	// Orders   []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	gorm.Model
}

type UpdateUserRequest struct {
	Nickname  string    `json:"nickname"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"omitempty,len=11"`
}
