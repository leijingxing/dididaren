package model

import (
	"time"

	"gorm.io/gorm"
)

// SecurityStaff 安保人员
type SecurityStaff struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Phone       string         `gorm:"size:20;not null" json:"phone"`
	IDCard      string         `gorm:"size:20;not null" json:"id_card"`
	Status      string         `gorm:"size:20;not null" json:"status"` // pending: 待审核 approved: 已通过 rejected: 已拒绝
	IsOnline    bool           `gorm:"default:false" json:"is_online"`
	LocationLat float64        `json:"location_lat"`
	LocationLng float64        `json:"location_lng"`
	OrderCount  int            `gorm:"default:0" json:"order_count"`
	Rating      float64        `gorm:"default:5.0" json:"rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SecurityStaff) TableName() string {
	return "security_staff"
}
