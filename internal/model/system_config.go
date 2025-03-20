package model

import (
	"time"
)

// SystemConfig 系统配置模型
type SystemConfig struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"uniqueIndex;not null"`
	Value     string    `json:"value" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateSystemConfigRequest 创建系统配置请求
type CreateSystemConfigRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Desc  string `json:"desc"`
}

// UpdateSystemConfigRequest 更新系统配置请求
type UpdateSystemConfigRequest struct {
	Value string `json:"value" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Desc  string `json:"desc"`
}
