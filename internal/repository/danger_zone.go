package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type DangerZoneRepository struct {
	db *gorm.DB
}

func NewDangerZoneRepository(db *gorm.DB) *DangerZoneRepository {
	return &DangerZoneRepository{db: db}
}

// CreateDangerZone 创建危险区域
func (r *DangerZoneRepository) CreateDangerZone(zone *model.DangerZone) error {
	return r.db.Create(zone).Error
}

// GetDangerZoneByID 根据ID获取危险区域
func (r *DangerZoneRepository) GetDangerZoneByID(id uint) (*model.DangerZone, error) {
	var zone model.DangerZone
	err := r.db.First(&zone, id).Error
	if err != nil {
		return nil, err
	}
	return &zone, nil
}

// ListDangerZones 获取危险区域列表
func (r *DangerZoneRepository) ListDangerZones(page, size int) ([]model.DangerZone, int64, error) {
	var zones []model.DangerZone
	var total int64

	err := r.db.Model(&model.DangerZone{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset((page - 1) * size).Limit(size).Find(&zones).Error
	if err != nil {
		return nil, 0, err
	}

	return zones, total, nil
}

// UpdateDangerZone 更新危险区域
func (r *DangerZoneRepository) UpdateDangerZone(zone *model.DangerZone) error {
	return r.db.Save(zone).Error
}

// DeleteDangerZone 删除危险区域
func (r *DangerZoneRepository) DeleteDangerZone(id uint) error {
	return r.db.Delete(&model.DangerZone{}, id).Error
}

// GetNearbyDangerZones 获取附近的危险区域
func (r *DangerZoneRepository) GetNearbyDangerZones(lat, lng float64) ([]model.DangerZone, error) {
	var zones []model.DangerZone
	// 这里使用简单的经纬度范围查询，实际项目中可能需要使用更复杂的空间查询
	// 例如使用 MySQL 的空间索引和 ST_Distance_Sphere 函数
	err := r.db.Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?",
		lat-0.1, lat+0.1, lng-0.1, lng+0.1).
		Find(&zones).Error
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// GetAllActiveZones 获取所有活跃的危险区域
func (r *DangerZoneRepository) GetAllActiveZones() ([]*model.DangerZone, error) {
	var zones []*model.DangerZone
	err := r.db.Where("status = ?", 1).Find(&zones).Error
	if err != nil {
		return nil, err
	}
	return zones, nil
}

// UpdateHeatLevel 更新危险区域的热度等级
func (r *DangerZoneRepository) UpdateHeatLevel(id uint, heatLevel int) error {
	return r.db.Model(&model.DangerZone{}).
		Where("id = ?", id).
		Update("heat_level", heatLevel).Error
}
