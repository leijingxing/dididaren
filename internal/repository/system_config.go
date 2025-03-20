package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type SystemConfigRepository struct {
	db *gorm.DB
}

func NewSystemConfigRepository(db *gorm.DB) *SystemConfigRepository {
	return &SystemConfigRepository{
		db: db,
	}
}

// Create 创建系统配置
func (r *SystemConfigRepository) Create(config *model.SystemConfig) error {
	return r.db.Create(config).Error
}

// GetByID 根据ID获取系统配置
func (r *SystemConfigRepository) GetByID(id uint) (*model.SystemConfig, error) {
	var config model.SystemConfig
	err := r.db.First(&config, id).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// List 获取系统配置列表
func (r *SystemConfigRepository) List(page, size int) ([]model.SystemConfig, int64, error) {
	var configs []model.SystemConfig
	var total int64

	// 获取总数
	if err := r.db.Model(&model.SystemConfig{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * size
	if err := r.db.Offset(offset).Limit(size).Find(&configs).Error; err != nil {
		return nil, 0, err
	}

	return configs, total, nil
}

// Update 更新系统配置
func (r *SystemConfigRepository) Update(config *model.SystemConfig) error {
	return r.db.Save(config).Error
}

// Delete 删除系统配置
func (r *SystemConfigRepository) Delete(id uint) error {
	return r.db.Delete(&model.SystemConfig{}, id).Error
}

// GetByKey 根据key获取配置
func (r *SystemConfigRepository) GetByKey(key string) (*model.SystemConfig, error) {
	var config model.SystemConfig
	err := r.db.Where("key = ?", key).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// GetValue 获取配置值
func (r *SystemConfigRepository) GetValue(key string) (string, error) {
	var config model.SystemConfig
	err := r.db.Where("key = ?", key).First(&config).Error
	if err != nil {
		return "", err
	}
	return config.Value, nil
}

// UpdateValue 更新配置值
func (r *SystemConfigRepository) UpdateValue(key string, value string) error {
	return r.db.Model(&model.SystemConfig{}).Where("key = ?", key).Update("value", value).Error
}
