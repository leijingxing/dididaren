package model

import (
	"time"

	"gorm.io/gorm"
)

// Staff 安保人员模型
type Staff struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	Name        string `gorm:"size:50;not null" json:"name"`
	Phone       string `gorm:"size:20;not null" json:"phone"`
	IDCard      string `gorm:"size:18;not null" json:"id_card"`
	Status      string `gorm:"size:20;not null;default:'pending'" json:"status"` // pending, active, inactive
	Rating      float64 `gorm:"default:0" json:"rating"`
	TotalOrders int    `gorm:"default:0" json:"total_orders"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	LastActive  time.Time `json:"last_active"`
}

// TableName 指定表名
func (Staff) TableName() string {
	return "staffs"
}

// CreateStaffRequest 创建安保人员请求
type CreateStaffRequest struct {
	Name   string `json:"name" binding:"required"`
	Phone  string `json:"phone" binding:"required"`
	IDCard string `json:"id_card" binding:"required"`
} 