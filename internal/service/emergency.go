package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"time"
)

type EmergencyService struct {
	repo *repository.EmergencyRepository
}

func NewEmergencyService(repo *repository.EmergencyRepository) *EmergencyService {
	return &EmergencyService{
		repo: repo,
	}
}

// Create 创建紧急事件
func (s *EmergencyService) Create(req *model.CreateEmergencyRequest) (*model.Emergency, error) {
	event := &model.Emergency{
		UserID:      req.UserID,
		EventType:   req.EventType,
		LocationLat: req.LocationLat,
		LocationLng: req.LocationLng,
		Address:     req.Address,
		Description: req.Description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Create(event)
}

// GetByID 获取事件详情
func (s *EmergencyService) GetByID(id string) (*model.Emergency, error) {
	return s.repo.GetByID(id)
}

// UpdateStatus 更新事件状态
func (s *EmergencyService) UpdateStatus(id string, status string) error {
	return s.repo.UpdateStatus(id, status)
}

// GetHistory 获取事件历史
func (s *EmergencyService) GetHistory(userID uint) ([]*model.Emergency, error) {
	return s.repo.GetHistory(userID)
}

// CreateHandlingRecord 创建处理记录
func (s *EmergencyService) CreateHandlingRecord(req *model.CreateHandlingRecordRequest) (*model.HandlingRecord, error) {
	record := &model.HandlingRecord{
		EventID:     req.EventID,
		StaffID:     req.StaffID,
		Action:      req.Action,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	return s.repo.CreateHandlingRecord(record)
}

// GetHandlingRecords 获取处理记录
func (s *EmergencyService) GetHandlingRecords(eventID string) ([]*model.HandlingRecord, error) {
	return s.repo.GetHandlingRecords(eventID)
}
