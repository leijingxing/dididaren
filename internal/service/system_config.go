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
	return &SystemConfigService{repo: repo}
}

// Create 创建系统配置
func (s *SystemConfigService) Create(req *model.CreateSystemConfigRequest) (*model.SystemConfig, error) {
	config := &model.SystemConfig{
		Key:   req.Key,
		Value: req.Value,
		Type:  req.Type,
		Desc:  req.Desc,
	}

	if err := s.repo.Create(config); err != nil {
		return nil, err
	}

	return config, nil
}

// GetByID 根据ID获取系统配置
func (s *SystemConfigService) GetByID(id uint) (*model.SystemConfig, error) {
	return s.repo.GetByID(id)
}

// List 获取系统配置列表
func (s *SystemConfigService) List(page, size int) ([]model.SystemConfig, int64, error) {
	return s.repo.List(page, size)
}

// Update 更新系统配置
func (s *SystemConfigService) Update(id uint, req *model.UpdateSystemConfigRequest) error {
	config, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	config.Value = req.Value
	config.Type = req.Type
	config.Desc = req.Desc

	return s.repo.Update(config)
}

// Delete 删除系统配置
func (s *SystemConfigService) Delete(id uint) error {
	return s.repo.Delete(id)
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

// GetValue 获取配置值
func (s *SystemConfigService) GetValue(key string) (string, error) {
	return s.repo.GetValue(key)
}

// UpdateValue 更新配置值
func (s *SystemConfigService) UpdateValue(key string, value string) error {
	return s.repo.UpdateValue(key, value)
}
