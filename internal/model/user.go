package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Avatar   string `gorm:"size:200" json:"avatar"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
	Status   int    `gorm:"default:1" json:"status"` // 1:正常 2:禁用
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

// UpdateProfileRequest 更新个人信息请求
type UpdateProfileRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
