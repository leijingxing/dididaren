package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
)

type SystemConfigService struct {
	repo *repository.SystemConfigRepository
}

func NewSystemConfigService(repo *repository.SystemConfigRepository) *SystemConfigService {
	return &SystemConfigService{
		repo: repo,
	}
}

// GetConfig 获取配置信息
func (s *SystemConfigService) GetConfig(key string) (*model.SystemConfig, error) {
	return s.repo.GetByKey(key)
}

// CreateConfig 创建配置
func (s *SystemConfigService) CreateConfig(config *model.SystemConfig) error {
	if config.ConfigKey == "" {
		return errors.New("配置键不能为空")
	}
	if config.ConfigValue == "" {
		return errors.New("配置值不能为空")
	}
	return s.repo.Create(config)
}

// UpdateConfig 更新配置
func (s *SystemConfigService) UpdateConfig(config *model.SystemConfig) error {
	if config.ConfigKey == "" {
		return errors.New("配置键不能为空")
	}
	if config.ConfigValue == "" {
		return errors.New("配置值不能为空")
	}
	return s.repo.Update(config)
}

// DeleteConfig 删除配置
func (s *SystemConfigService) DeleteConfig(key string) error {
	return s.repo.Delete(key)
}

// GetAllConfigs 获取所有配置
func (s *SystemConfigService) GetAllConfigs() ([]*model.SystemConfig, error) {
	return s.repo.GetAll()
}

// GetConfigValue 获取配置值
func (s *SystemConfigService) GetConfigValue(key string) (string, error) {
	return s.repo.GetValueByKey(key)
}

// UpdateConfigValue 更新配置值
func (s *SystemConfigService) UpdateConfigValue(key string, value string) error {
	if value == "" {
		return errors.New("配置值不能为空")
	}
	return s.repo.UpdateValue(key, value)
}
