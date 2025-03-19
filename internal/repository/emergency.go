package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type EmergencyRepository struct {
	db *gorm.DB
}

func NewEmergencyRepository(db *gorm.DB) *EmergencyRepository {
	return &EmergencyRepository{
		db: db,
	}
}

// Create 创建紧急事件
func (r *EmergencyRepository) Create(event *model.Emergency) (*model.Emergency, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

// GetByID 获取事件详情
func (r *EmergencyRepository) GetByID(id string) (*model.Emergency, error) {
	var event model.Emergency
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// UpdateStatus 更新事件状态
func (r *EmergencyRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Emergency{}).Where("id = ?", id).Update("status", status).Error
}

// GetHistory 获取事件历史
func (r *EmergencyRepository) GetHistory(userID uint) ([]*model.Emergency, error) {
	var events []*model.Emergency
	if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

// CreateHandlingRecord 创建处理记录
func (r *EmergencyRepository) CreateHandlingRecord(record *model.HandlingRecord) (*model.HandlingRecord, error) {
	if err := r.db.Create(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

// GetHandlingRecords 获取处理记录
func (r *EmergencyRepository) GetHandlingRecords(eventID string) ([]*model.HandlingRecord, error) {
	var records []*model.HandlingRecord
	if err := r.db.Where("event_id = ?", eventID).Order("created_at desc").Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
