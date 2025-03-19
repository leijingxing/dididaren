package model

import (
	"time"

	"gorm.io/gorm"
)

// Emergency 紧急事件模型
type Emergency struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"user_id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Location    string    `gorm:"size:200;not null" json:"location"`
	Latitude    float64   `gorm:"not null" json:"latitude"`
	Longitude   float64   `gorm:"not null" json:"longitude"`
	Level       string    `gorm:"size:20;not null" json:"level"` // low, medium, high
	Status      string    `gorm:"size:20;not null;default:'pending'" json:"status"` // pending, processing, completed, cancelled
	StaffID     uint      `json:"staff_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// TableName 指定表名
func (Emergency) TableName() string {
	return "emergencies"
}

// CreateEmergencyRequest 创建紧急事件请求
type CreateEmergencyRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Location    string  `json:"location" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
	Level       string  `json:"level" binding:"required,oneof=low medium high"`
}

// EmergencyContact 紧急联系人模型
type EmergencyContact struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	Name        string `gorm:"size:50;not null" json:"name"`
	Phone       string `gorm:"size:20;not null" json:"phone"`
	Relation    string `gorm:"size:50" json:"relation"`
	IsDefault   bool   `gorm:"default:false" json:"is_default"`
}

// TableName 指定表名
func (EmergencyContact) TableName() string {
	return "emergency_contacts"
}

// HandlingRecord 处理记录模型
type HandlingRecord struct {
	gorm.Model
	EmergencyID uint      `gorm:"not null" json:"emergency_id"`
	StaffID     uint      `gorm:"not null" json:"staff_id"`
	Action      string    `gorm:"size:50;not null" json:"action"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"size:20;not null" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// TableName 指定表名
func (HandlingRecord) TableName() string {
	return "handling_records"
}

// CreateHandlingRecordRequest 创建处理记录请求
type CreateHandlingRecordRequest struct {
	Action      string `json:"action" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required,oneof=pending processing completed cancelled"`
}
