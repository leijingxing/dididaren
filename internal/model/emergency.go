package model

import (
	"time"

	"gorm.io/gorm"
)

// Emergency 紧急事件
type Emergency struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	EventType   int8           `gorm:"not null" json:"event_type"`    // 1:普通求助 2:家庭暴力 3:医疗急救 4:诈骗干预
	Status      string         `gorm:"default:pending" json:"status"` // pending: 待处理 processing: 处理中 completed: 已完成 cancelled: 已取消
	LocationLat float64        `gorm:"not null" json:"location_lat"`
	LocationLng float64        `gorm:"not null" json:"location_lng"`
	Address     string         `gorm:"size:255" json:"address"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// EmergencyContact 紧急联系人
type EmergencyContact struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	Name         string         `gorm:"size:50;not null" json:"name"`
	Phone        string         `gorm:"size:20;not null" json:"phone"`
	Relationship string         `gorm:"size:20" json:"relationship"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// HandlingRecord 事件处理记录
type HandlingRecord struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	EventID     uint           `gorm:"not null" json:"event_id"`
	StaffID     uint           `gorm:"not null" json:"staff_id"`
	Action      string         `gorm:"not null" json:"action"` // accept: 接单 arrive: 到达 complete: 完成
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// DangerZone 危险区域
type DangerZone struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	LocationLat float64        `gorm:"not null" json:"location_lat"`
	LocationLng float64        `gorm:"not null" json:"location_lng"`
	Radius      int            `json:"radius"`                       // 危险区域半径(米)
	DangerLevel int8           `gorm:"not null" json:"danger_level"` // 1:低 2:中 3:高
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// CreateEmergencyRequest 创建紧急事件请求
type CreateEmergencyRequest struct {
	UserID      uint    `json:"user_id"`
	EventType   int8    `json:"event_type" binding:"required"`
	LocationLat float64 `json:"location_lat" binding:"required"`
	LocationLng float64 `json:"location_lng" binding:"required"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
}

// CreateHandlingRecordRequest 创建处理记录请求
type CreateHandlingRecordRequest struct {
	EventID     uint   `json:"event_id" binding:"required"`
	StaffID     uint   `json:"staff_id" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Description string `json:"description"`
}

func (Emergency) TableName() string {
	return "emergencies"
}

func (EmergencyContact) TableName() string {
	return "emergency_contacts"
}

func (HandlingRecord) TableName() string {
	return "handling_records"
}

func (DangerZone) TableName() string {
	return "danger_zones"
}
