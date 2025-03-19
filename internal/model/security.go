package model

import (
	"time"

	"gorm.io/gorm"
)

// SecurityStaff 安保人员
type SecurityStaff struct {
	ID                  uint           `gorm:"primarykey" json:"id"`
	UserID              uint           `gorm:"not null" json:"user_id"`
	CompanyName         string         `gorm:"size:100" json:"company_name"`
	LicenseNumber       string         `gorm:"size:50" json:"license_number"`
	CertificationStatus int8           `gorm:"default:0" json:"certification_status"` // 0:待审核 1:已认证 2:已拒绝
	Rating              float64        `gorm:"default:5.0" json:"rating"`
	TotalOrders         int            `gorm:"default:0" json:"total_orders"`
	OnlineStatus        int8           `gorm:"default:0" json:"online_status"` // 0:离线 1:在线
	CurrentLocationLat  float64        `json:"current_location_lat"`
	CurrentLocationLng  float64        `json:"current_location_lng"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}

// Rating 评价
type Rating struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	EventID   uint           `gorm:"not null" json:"event_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	StaffID   uint           `gorm:"not null" json:"staff_id"`
	Rating    int8           `gorm:"not null" json:"rating"` // 1-5星
	Comment   string         `gorm:"type:text" json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SecurityStaff) TableName() string {
	return "security_staff"
}

func (Rating) TableName() string {
	return "ratings"
}
