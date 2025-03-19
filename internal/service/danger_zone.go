package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
)

type DangerZoneService struct {
	repo *repository.DangerZoneRepository
}

func NewDangerZoneService(repo *repository.DangerZoneRepository) *DangerZoneService {
	return &DangerZoneService{
		repo: repo,
	}
}

// CreateDangerZone 创建危险区域
func (s *DangerZoneService) CreateDangerZone(zone *model.DangerZone) error {
	if zone.Latitude < -90 || zone.Latitude > 90 {
		return errors.New("纬度范围无效")
	}
	if zone.Longitude < -180 || zone.Longitude > 180 {
		return errors.New("经度范围无效")
	}
	if zone.Radius <= 0 {
		return errors.New("半径必须大于0")
	}
	return s.repo.Create(zone)
}

// GetDangerZone 获取危险区域信息
func (s *DangerZoneService) GetDangerZone(id uint) (*model.DangerZone, error) {
	return s.repo.GetByID(id)
}

// UpdateDangerZone 更新危险区域信息
func (s *DangerZoneService) UpdateDangerZone(zone *model.DangerZone) error {
	if zone.Latitude < -90 || zone.Latitude > 90 {
		return errors.New("纬度范围无效")
	}
	if zone.Longitude < -180 || zone.Longitude > 180 {
		return errors.New("经度范围无效")
	}
	if zone.Radius <= 0 {
		return errors.New("半径必须大于0")
	}
	return s.repo.Update(zone)
}

// DeleteDangerZone 删除危险区域
func (s *DangerZoneService) DeleteDangerZone(id uint) error {
	return s.repo.Delete(id)
}

// GetNearbyZones 获取附近的危险区域
func (s *DangerZoneService) GetNearbyZones(latitude, longitude float64, radius float64) ([]*model.DangerZone, error) {
	if latitude < -90 || latitude > 90 {
		return nil, errors.New("纬度范围无效")
	}
	if longitude < -180 || longitude > 180 {
		return nil, errors.New("经度范围无效")
	}
	if radius <= 0 {
		return nil, errors.New("搜索半径必须大于0")
	}
	return s.repo.GetNearbyZones(latitude, longitude, radius)
}

// GetAllActiveZones 获取所有活跃的危险区域
func (s *DangerZoneService) GetAllActiveZones() ([]*model.DangerZone, error) {
	return s.repo.GetAllActiveZones()
}

// UpdateHeatLevel 更新危险区域的热度等级
func (s *DangerZoneService) UpdateHeatLevel(id uint, heatLevel int) error {
	if heatLevel < 0 || heatLevel > 5 {
		return errors.New("热度等级必须在0-5之间")
	}
	return s.repo.UpdateHeatLevel(id, heatLevel)
}
