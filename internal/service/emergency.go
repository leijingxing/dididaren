package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
	"time"
)

type EmergencyService struct {
	repo *repository.EmergencyRepository
}

func NewEmergencyService(repo *repository.EmergencyRepository) *EmergencyService {
	return &EmergencyService{repo: repo}
}

// Create 创建紧急事件
func (s *EmergencyService) Create(userID uint, req *model.CreateEmergencyRequest) (*model.Emergency, error) {
	emergency := &model.Emergency{
		UserID:      userID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Status:      1, // 待处理
	}

	if err := s.repo.Create(emergency); err != nil {
		return nil, err
	}

	return emergency, nil
}

// GetByID 获取紧急事件详情
func (s *EmergencyService) GetByID(id uint) (*model.Emergency, error) {
	return s.repo.GetByID(id)
}

// List 获取紧急事件列表
func (s *EmergencyService) List(page, size int) ([]model.Emergency, int64, error) {
	return s.repo.List(page, size)
}

// Update 更新紧急事件
func (s *EmergencyService) Update(id uint, req *model.UpdateEmergencyRequest) (*model.Emergency, error) {
	emergency, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		emergency.Title = req.Title
	}
	if req.Description != "" {
		emergency.Description = req.Description
	}
	if req.Location != "" {
		emergency.Location = req.Location
	}
	if req.Latitude != 0 {
		emergency.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		emergency.Longitude = req.Longitude
	}

	if err := s.repo.Update(emergency); err != nil {
		return nil, err
	}

	return emergency, nil
}

// Delete 删除紧急事件
func (s *EmergencyService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// CreateHandlingRecord 创建处理记录
func (s *EmergencyService) CreateHandlingRecord(emergencyID uint, req *model.CreateHandlingRecordRequest) (*model.HandlingRecord, error) {
	record := &model.HandlingRecord{
		EmergencyID: emergencyID,
		Action:      req.Action,
		Description: req.Description,
	}

	if err := s.repo.CreateHandlingRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// ListHandlingRecords 获取处理记录列表
func (s *EmergencyService) ListHandlingRecords(emergencyID uint) ([]model.HandlingRecord, error) {
	return s.repo.ListHandlingRecords(emergencyID)
}

// UpdateStatus 更新紧急事件状态
func (s *EmergencyService) UpdateStatus(id uint, status int) error {
	emergency, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	emergency.Status = status
	return s.repo.Update(emergency)
}

func (s *EmergencyService) CreateEmergency(userID uint, req *model.CreateEmergencyRequest) (*model.Emergency, error) {
	emergency := &model.Emergency{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Level:       req.Level,
		Status:      "pending",
		StartTime:   time.Now(),
	}

	if err := s.repo.CreateEmergency(emergency); err != nil {
		return nil, err
	}

	return emergency, nil
}

func (s *EmergencyService) GetEmergencyByID(id uint) (*model.Emergency, error) {
	return s.repo.GetEmergencyByID(id)
}

func (s *EmergencyService) ListEmergencies(page, size int, status string) ([]model.Emergency, int64, error) {
	return s.repo.ListEmergencies(page, size, status)
}

func (s *EmergencyService) UpdateEmergencyStatus(id uint, status string) error {
	if status != "pending" && status != "processing" && status != "completed" && status != "cancelled" {
		return errors.New("无效的状态")
	}
	return s.repo.UpdateEmergencyStatus(id, status)
}

func (s *EmergencyService) AssignStaff(emergencyID, staffID uint) error {
	emergency, err := s.repo.GetEmergencyByID(emergencyID)
	if err != nil {
		return err
	}

	if emergency.Status != "pending" {
		return errors.New("该事件已被处理")
	}

	emergency.StaffID = staffID
	emergency.Status = "processing"
	return s.repo.UpdateEmergency(emergency)
}

func (s *EmergencyService) CompleteEmergency(emergencyID uint) error {
	emergency, err := s.repo.GetEmergencyByID(emergencyID)
	if err != nil {
		return err
	}

	if emergency.Status != "processing" {
		return errors.New("该事件状态不正确")
	}

	emergency.Status = "completed"
	emergency.EndTime = time.Now()
	return s.repo.UpdateEmergency(emergency)
}
