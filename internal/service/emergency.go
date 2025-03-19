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

func (s *EmergencyService) CreateHandlingRecord(req *model.CreateHandlingRecordRequest, emergencyID, staffID uint) (*model.HandlingRecord, error) {
	record := &model.HandlingRecord{
		EmergencyID: emergencyID,
		StaffID:     staffID,
		Action:      req.Action,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.CreateHandlingRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (s *EmergencyService) ListHandlingRecords(emergencyID uint) ([]model.HandlingRecord, error) {
	return s.repo.ListHandlingRecords(emergencyID)
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
