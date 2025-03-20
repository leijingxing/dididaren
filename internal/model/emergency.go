package model

import (
	"time"

	"gorm.io/gorm"
)

// Emergency 紧急事件
type Emergency struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Emergency) TableName() string {
	return "emergencies"
}

// CreateEmergencyRequest 创建紧急事件请求
type CreateEmergencyRequest struct {
	Type        string  `json:"type" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Location    string  `json:"location" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" binding:"required"`
}

// UpdateEmergencyRequest 更新紧急事件请求
type UpdateEmergencyRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// UpdateEmergencyStatusRequest 更新紧急事件状态请求
type UpdateEmergencyStatusRequest struct {
	Status int `json:"status" binding:"required"`
}

// EmergencyContact 紧急联系人模型
type EmergencyContact struct {
	gorm.Model
	UserID    uint   `gorm:"not null" json:"user_id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Phone     string `gorm:"size:20;not null" json:"phone"`
	Relation  string `gorm:"size:50" json:"relation"`
	IsDefault bool   `gorm:"default:false" json:"is_default"`
}

// TableName 指定表名
func (EmergencyContact) TableName() string {
	return "emergency_contacts"
}

// HandlingRecord 处理记录
type HandlingRecord struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	EmergencyID uint      `json:"emergency_id"`
	StaffID     uint      `json:"staff_id"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName 指定表名
func (HandlingRecord) TableName() string {
	return "handling_records"
}

// CreateHandlingRecordRequest 创建处理记录请求
type CreateHandlingRecordRequest struct {
	Action      string `json:"action" binding:"required"`
	Description string `json:"description" binding:"required"`
}
