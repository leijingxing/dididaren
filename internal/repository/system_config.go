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

// GetByKey 根据配置键获取配置
func (r *SystemConfigRepository) GetByKey(key string) (*model.SystemConfig, error) {
	var config model.SystemConfig
	err := r.db.Where("config_key = ?", key).First(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}

// Create 创建配置
func (r *SystemConfigRepository) Create(config *model.SystemConfig) error {
	return r.db.Create(config).Error
}

// Update 更新配置
func (r *SystemConfigRepository) Update(config *model.SystemConfig) error {
	return r.db.Save(config).Error
}

// Delete 删除配置
func (r *SystemConfigRepository) Delete(key string) error {
	return r.db.Where("config_key = ?", key).Delete(&model.SystemConfig{}).Error
}

// GetAll 获取所有配置
func (r *SystemConfigRepository) GetAll() ([]*model.SystemConfig, error) {
	var configs []*model.SystemConfig
	err := r.db.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}

// GetValueByKey 获取配置值
func (r *SystemConfigRepository) GetValueByKey(key string) (string, error) {
	var config model.SystemConfig
	err := r.db.Where("config_key = ?", key).First(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return config.ConfigValue, nil
}

// UpdateValue 更新配置值
func (r *SystemConfigRepository) UpdateValue(key string, value string) error {
	return r.db.Model(&model.SystemConfig{}).
		Where("config_key = ?", key).
		Update("config_value", value).Error
}
