package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
)

type DangerZoneService struct {
	repo *repository.DangerZoneRepository
}

func NewDangerZoneService(repo *repository.DangerZoneRepository) *DangerZoneService {
	return &DangerZoneService{repo: repo}
}

// CreateDangerZone 创建危险区域
func (s *DangerZoneService) CreateDangerZone(req *model.CreateDangerZoneRequest) error {
	zone := &model.DangerZone{
		Name:        req.Name,
		Description: req.Description,
		Level:       req.Level,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Radius:      req.Radius,
	}

	return s.repo.CreateDangerZone(zone)
}

// GetDangerZoneByID 根据ID获取危险区域
func (s *DangerZoneService) GetDangerZoneByID(id uint) (*model.DangerZone, error) {
	return s.repo.GetDangerZoneByID(id)
}

// ListDangerZones 获取危险区域列表
func (s *DangerZoneService) ListDangerZones(page, size int) ([]model.DangerZone, int64, error) {
	return s.repo.ListDangerZones(page, size)
}

// UpdateDangerZone 更新危险区域
func (s *DangerZoneService) UpdateDangerZone(id uint, req *model.UpdateDangerZoneRequest) error {
	zone, err := s.repo.GetDangerZoneByID(id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		zone.Name = req.Name
	}
	if req.Description != "" {
		zone.Description = req.Description
	}
	if req.Level != "" {
		zone.Level = req.Level
	}
	if req.Latitude != 0 {
		zone.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		zone.Longitude = req.Longitude
	}
	if req.Radius != 0 {
		zone.Radius = req.Radius
	}

	return s.repo.UpdateDangerZone(zone)
}

// DeleteDangerZone 删除危险区域
func (s *DangerZoneService) DeleteDangerZone(id uint) error {
	return s.repo.DeleteDangerZone(id)
}

// CheckLocationInDangerZone 检查位置是否在危险区域内
func (s *DangerZoneService) CheckLocationInDangerZone(lat, lng float64) (bool, []model.DangerZone, error) {
	zones, err := s.repo.GetNearbyDangerZones(lat, lng)
	if err != nil {
		return false, nil, err
	}

	if len(zones) > 0 {
		return true, zones, nil
	}

	return false, nil, nil
}
