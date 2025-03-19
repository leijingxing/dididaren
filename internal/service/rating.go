package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
)

type RatingService struct {
	repo *repository.RatingRepository
}

func NewRatingService(repo *repository.RatingRepository) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

// CreateRating 创建评价
func (s *RatingService) CreateRating(rating *model.Rating) error {
	if rating.Score < 1 || rating.Score > 5 {
		return errors.New("评分必须在1-5之间")
	}
	if rating.EventID == 0 {
		return errors.New("事件ID不能为空")
	}
	if rating.StaffID == 0 {
		return errors.New("安保人员ID不能为空")
	}
	return s.repo.Create(rating)
}

// GetRating 获取评价信息
func (s *RatingService) GetRating(id uint) (*model.Rating, error) {
	return s.repo.GetByID(id)
}

// GetEventRating 获取事件相关的评价
func (s *RatingService) GetEventRating(eventID uint) (*model.Rating, error) {
	return s.repo.GetByEventID(eventID)
}

// GetStaffRatings 获取安保人员的所有评价
func (s *RatingService) GetStaffRatings(staffID uint) ([]*model.Rating, error) {
	return s.repo.GetStaffRatings(staffID)
}

// GetStaffAverageRating 获取安保人员的平均评分
func (s *RatingService) GetStaffAverageRating(staffID uint) (float64, error) {
	return s.repo.CalculateStaffAverageRating(staffID)
}

// UpdateRating 更新评价
func (s *RatingService) UpdateRating(rating *model.Rating) error {
	if rating.Score < 1 || rating.Score > 5 {
		return errors.New("评分必须在1-5之间")
	}
	return s.repo.Update(rating)
}

// DeleteRating 删除评价
func (s *RatingService) DeleteRating(id uint) error {
	return s.repo.Delete(id)
}
