package model

import (
	"time"
)

// Rating 评价模型
type Rating struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StaffID   uint      `json:"staff_id"`
	UserID    uint      `json:"user_id"`
	Score     float32   `json:"score"`
	Comment   string    `json:"comment"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateRatingRequest 创建评价请求
type CreateRatingRequest struct {
	StaffID  uint    `json:"staff_id" binding:"required"`
	UserID   uint    `json:"user_id" binding:"required"`
	Score    float32 `json:"score" binding:"required,min=0,max=5"`
	Comment  string  `json:"comment" binding:"required,min=1,max=500"`
	IsPublic bool    `json:"is_public"`
}
