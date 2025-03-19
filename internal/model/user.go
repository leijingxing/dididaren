package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Phone     string    `gorm:"uniqueIndex;size:11;not null" json:"phone"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Name      string    `gorm:"size:32" json:"name"`
	Avatar    string    `gorm:"size:256" json:"avatar"`
	LastLogin time.Time `json:"last_login"`
	IsAdmin   bool      `gorm:"default:false" json:"is_admin"`
	Status    int       `gorm:"default:1" json:"status"` // 1:正常 2:禁用
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Name     string `json:"name" binding:"required,max=32"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest 更新用户信息请求
type UpdateProfileRequest struct {
	Name   string `json:"name" binding:"required,max=32"`
	Avatar string `json:"avatar" binding:"omitempty,max=256"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
