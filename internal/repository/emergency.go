package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type EmergencyRepository struct {
	db *gorm.DB
}

func NewEmergencyRepository(db *gorm.DB) *EmergencyRepository {
	return &EmergencyRepository{db: db}
}

// Create 创建紧急事件
func (r *EmergencyRepository) Create(emergency *model.Emergency) error {
	return r.db.Create(emergency).Error
}

// GetByID 获取紧急事件详情
func (r *EmergencyRepository) GetByID(id uint) (*model.Emergency, error) {
	var emergency model.Emergency
	err := r.db.First(&emergency, id).Error
	if err != nil {
		return nil, err
	}
	return &emergency, nil
}

// List 获取紧急事件列表
func (r *EmergencyRepository) List(page, size int) ([]model.Emergency, int64, error) {
	var emergencies []model.Emergency
	var total int64

	err := r.db.Model(&model.Emergency{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset((page - 1) * size).Limit(size).Find(&emergencies).Error
	if err != nil {
		return nil, 0, err
	}

	return emergencies, total, nil
}

// Update 更新紧急事件
func (r *EmergencyRepository) Update(emergency *model.Emergency) error {
	return r.db.Save(emergency).Error
}

// Delete 删除紧急事件
func (r *EmergencyRepository) Delete(id uint) error {
	return r.db.Delete(&model.Emergency{}, id).Error
}

// CreateHandlingRecord 创建处理记录
func (r *EmergencyRepository) CreateHandlingRecord(record *model.HandlingRecord) error {
	return r.db.Create(record).Error
}

// ListHandlingRecords 获取处理记录列表
func (r *EmergencyRepository) ListHandlingRecords(emergencyID uint) ([]model.HandlingRecord, error) {
	var records []model.HandlingRecord
	err := r.db.Where("emergency_id = ?", emergencyID).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}
