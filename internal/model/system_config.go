package model

import (
	"gorm.io/gorm"
)

// SystemConfig 系统配置模型
type SystemConfig struct {
	gorm.Model
	Key    string `gorm:"size:50;not null;unique" json:"key"`
	Value  string `gorm:"type:text" json:"value"`
	Type   string `gorm:"size:20;not null" json:"type"` // string, number, boolean, json
	Remark string `gorm:"size:200" json:"remark"`
}

// TableName 指定表名
func (SystemConfig) TableName() string {
	return "system_configs"
}

// CreateConfigRequest 创建配置请求
type CreateConfigRequest struct {
	Key    string `json:"key" binding:"required"`
	Value  string `json:"value" binding:"required"`
	Type   string `json:"type" binding:"required,oneof=string number boolean json"`
	Remark string `json:"remark"`
}

// UpdateConfigRequest 更新配置请求
type UpdateConfigRequest struct {
	Value  string `json:"value" binding:"required"`
	Type   string `json:"type" binding:"required,oneof=string number boolean json"`
	Remark string `json:"remark"`
}
