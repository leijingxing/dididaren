package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type DangerZoneRepository struct {
	db *gorm.DB
}

func NewDangerZoneRepository(db *gorm.DB) *DangerZoneRepository {
	return &DangerZoneRepository{
		db: db,
	}
}

// Create 创建危险区域
func (r *DangerZoneRepository) Create(zone *model.DangerZone) error {
	return r.db.Create(zone).Error
}

// GetByID 根据ID获取危险区域
func (r *DangerZoneRepository) GetByID(id uint) (*model.DangerZone, error) {
	var zone model.DangerZone
	err := r.db.First(&zone, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &zone, nil
}

// Update 更新危险区域信息
func (r *DangerZoneRepository) Update(zone *model.DangerZone) error {
	return r.db.Save(zone).Error
}

// Delete 删除危险区域
func (r *DangerZoneRepository) Delete(id uint) error {
	return r.db.Delete(&model.DangerZone{}, id).Error
}

// GetNearbyZones 获取附近的危险区域
func (r *DangerZoneRepository) GetNearbyZones(latitude, longitude float64, radius float64) ([]*model.DangerZone, error) {
	var zones []*model.DangerZone
	err := r.db.Where("latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?",
		latitude-radius, latitude+radius,
		longitude-radius, longitude+radius,
	).Find(&zones).Error
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
