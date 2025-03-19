package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"dididaren/pkg/errors"
)

type SystemConfigService struct {
	repo *repository.SystemConfigRepository
}

func NewSystemConfigService(repo *repository.SystemConfigRepository) *SystemConfigService {
	return &SystemConfigService{
		repo: repo,
	}
}

// Create 创建配置
func (s *SystemConfigService) Create(key, value, typ, remark string) error {
	config := &model.SystemConfig{
		Key:    key,
		Value:  value,
		Type:   typ,
		Remark: remark,
	}

	return s.repo.Create(config)
}

// GetByKey 根据key获取配置
func (s *SystemConfigService) GetByKey(key string) (*model.SystemConfig, error) {
	config, err := s.repo.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errors.ErrConfigNotFound
	}
	return config, nil
}

// Update 更新配置
func (s *SystemConfigService) Update(key, value, typ, remark string) error {
	config, err := s.repo.GetByKey(key)
	if err != nil {
		return err
	}
	if config == nil {
		return errors.ErrConfigNotFound
	}

	config.Value = value
	config.Type = typ
	config.Remark = remark

	return s.repo.Update(config)
}

// Delete 删除配置
func (s *SystemConfigService) Delete(key string) error {
	return s.repo.Delete(key)
}

// List 获取配置列表
func (s *SystemConfigService) List() ([]*model.SystemConfig, error) {
	return s.repo.List()
}

// GetValue 获取配置值
func (s *SystemConfigService) GetValue(key string) (string, error) {
	return s.repo.GetValue(key)
}

// UpdateValue 更新配置值
func (s *SystemConfigService) UpdateValue(key string, value string) error {
	return s.repo.UpdateValue(key, value)
}
