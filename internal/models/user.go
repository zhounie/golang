package models

import (
	"time"
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
	// // @swaggo.ignore
	// gorm.Model
	// 用户名
	Username string `gorm:"size:50;uniqueIndex;not null" json:"username"`
	// 邮箱
	Email string `gorm:"size:100;uniqueIndex;not null" json:"email"`
	// 密码
	Password string `gorm:"size:255;not null" json:"password"`
	// 手机号
	Phone string `gorm:"size:20" json:"phone"`
	// 头像
	Avatar string `gorm:"size:255" json:"avatar"`
	// 角色
	Role string `gorm:"size:20;default:'customer'" json:"role"`
	// 状态
	Status string `gorm:"size:20;default:'active'" json:"status"`

	// 昵称
	Nickname string `gorm:"size:100" json:"nickname"`
	// 性别
	Gender string `gorm:"size:10" json:"gender"`
	// 生日
	Birthdate *time.Time `json:"birthdate"`

	// 最近登录时间
	LastLoginAt *time.Time `json:last_login_at`

	// Products []Product `gorm:"foreignKey:UserID" json:"pruducts,omitempty"`
	// Orders   []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`

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
