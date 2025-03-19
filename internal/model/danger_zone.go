package model

import (
	"gorm.io/gorm"
)

// DangerZone 危险区域模型
type DangerZone struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Level       string  `gorm:"size:20;not null" json:"level"` // low, medium, high
	Latitude    float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	Radius      float64 `gorm:"not null" json:"radius"` // 半径（米）
	HeatLevel   int     `gorm:"default:0" json:"heat_level"` // 热度等级
	IsActive    bool    `gorm:"default:true" json:"is_active"`
}

// TableName 指定表名
func (DangerZone) TableName() string {
	return "danger_zones"
}

// CreateDangerZoneRequest 创建危险区域请求
type CreateDangerZoneRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Level       string  `json:"level" binding:"required,oneof=low medium high"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	Radius      float64 `json:"radius" binding:"required,min=0"`
}

// UpdateDangerZoneRequest 更新危险区域请求
type UpdateDangerZoneRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Level       string  `json:"level" binding:"omitempty,oneof=low medium high"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Radius      float64 `json:"radius" binding:"omitempty,min=0"`
	IsActive    bool    `json:"is_active"`
}
