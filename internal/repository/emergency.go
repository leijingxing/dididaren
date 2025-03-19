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

// CreateEmergency 创建紧急事件
func (r *EmergencyRepository) CreateEmergency(emergency *model.Emergency) error {
	return r.db.Create(emergency).Error
}

// GetEmergencyByID 根据ID获取紧急事件
func (r *EmergencyRepository) GetEmergencyByID(id uint) (*model.Emergency, error) {
	var emergency model.Emergency
	err := r.db.First(&emergency, id).Error
	if err != nil {
		return nil, err
	}
	return &emergency, nil
}

// ListEmergencies 获取紧急事件列表
func (r *EmergencyRepository) ListEmergencies(page, size int, status string) ([]model.Emergency, int64, error) {
	var emergencies []model.Emergency
	var total int64

	query := r.db.Model(&model.Emergency{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * size).Limit(size).Find(&emergencies).Error
	if err != nil {
		return nil, 0, err
	}

	return emergencies, total, nil
}

// UpdateEmergencyStatus 更新紧急事件状态
func (r *EmergencyRepository) UpdateEmergencyStatus(id uint, status string) error {
	return r.db.Model(&model.Emergency{}).Where("id = ?", id).Update("status", status).Error
}

// UpdateEmergency 更新紧急事件
func (r *EmergencyRepository) UpdateEmergency(emergency *model.Emergency) error {
	return r.db.Save(emergency).Error
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
