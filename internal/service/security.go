package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
)

type SecurityService struct {
	repo *repository.SecurityRepository
}

func NewSecurityService(repo *repository.SecurityRepository) *SecurityService {
	return &SecurityService{repo: repo}
}

func (s *SecurityService) CreateStaff(req *model.CreateStaffRequest) (*model.Staff, error) {
	staff := &model.Staff{
		Name:   req.Name,
		Phone:  req.Phone,
		IDCard: req.IDCard,
	}

	if err := s.repo.CreateStaff(staff); err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *SecurityService) GetStaffByID(id uint) (*model.Staff, error) {
	return s.repo.GetStaffByID(id)
}

func (s *SecurityService) ListStaffs(page, size int, status string) ([]model.Staff, int64, error) {
	return s.repo.ListStaffs(page, size, status)
}

func (s *SecurityService) UpdateStaffStatus(id uint, status string) error {
	if status != "pending" && status != "active" && status != "inactive" {
		return errors.New("无效的状态")
	}
	return s.repo.UpdateStaffStatus(id, status)
}

func (s *SecurityService) CreateRating(req *model.CreateRatingRequest) (*model.Rating, error) {
	rating := &model.Rating{
		StaffID:  req.StaffID,
		UserID:   req.UserID,
		Score:    req.Score,
		Comment:  req.Comment,
		IsPublic: req.IsPublic,
	}

	if err := s.repo.CreateRating(rating); err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *SecurityService) ListRatings(staffID uint) ([]model.Rating, error) {
	return s.repo.ListRatings(staffID)
}

func (s *SecurityService) UpdateLocation(staffID uint, lat, lng float64) error {
	return s.repo.UpdateLocation(staffID, lat, lng)
}

func (s *SecurityService) GetStaffInfo(userID uint) (*model.Staff, error) {
	return s.repo.GetStaffByUserID(userID)
}

func (s *SecurityService) ApplySecurityStaff(userID uint, name, phone, idCard string) error {
	// 检查是否已经是安保人员
	existingStaff, err := s.repo.GetStaffByUserID(userID)
	if err == nil && existingStaff != nil {
		return errors.New("您已经是安保人员")
	}

	staff := &model.Staff{
		UserID: userID,
		Name:   name,
		Phone:  phone,
		IDCard: idCard,
	}

	if err := s.repo.CreateStaff(staff); err != nil {
		return err
	}
	return nil
}

func (s *SecurityService) AcceptEvent(staffID uint, eventID uint) error {
	// TODO: 实现接单逻辑
	return nil
}

func (s *SecurityService) CompleteEvent(staffID uint, eventID uint) error {
	// TODO: 实现完成订单逻辑
	return nil
}
