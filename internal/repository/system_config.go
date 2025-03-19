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

// Create 创建配置
func (r *SystemConfigRepository) Create(config *model.SystemConfig) error {
	return r.db.Create(config).Error
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

// Update 更新配置
func (r *SystemConfigRepository) Update(config *model.SystemConfig) error {
	return r.db.Save(config).Error
}

// Delete 删除配置
func (r *SystemConfigRepository) Delete(key string) error {
	return r.db.Where("key = ?", key).Delete(&model.SystemConfig{}).Error
}

// List 获取配置列表
func (r *SystemConfigRepository) List() ([]*model.SystemConfig, error) {
	var configs []*model.SystemConfig
	err := r.db.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
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
