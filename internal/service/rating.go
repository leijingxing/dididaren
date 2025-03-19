package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
)

type RatingService struct {
	repo *repository.RatingRepository
}

func NewRatingService(repo *repository.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

// CreateRating 创建评价
func (s *RatingService) CreateRating(req *model.CreateRatingRequest) (*model.Rating, error) {
	rating := &model.Rating{
		StaffID:  req.StaffID,
		UserID:   req.UserID,
		Score:    req.Score,
		Comment:  req.Comment,
		IsPublic: req.IsPublic,
	}
	return s.repo.CreateRating(rating)
}

// GetRatingByID 根据ID获取评价
func (s *RatingService) GetRatingByID(id uint) (*model.Rating, error) {
	return s.repo.GetRatingByID(id)
}

// ListRatings 获取安保人员的评价列表
func (s *RatingService) ListRatings(staffID uint) ([]*model.Rating, error) {
	return s.repo.ListRatings(staffID)
}

// UpdateRating 更新评价
func (s *RatingService) UpdateRating(id uint, req *model.CreateRatingRequest) error {
	rating, err := s.repo.GetRatingByID(id)
	if err != nil {
		return err
	}

	rating.Score = req.Score
	rating.Comment = req.Comment
	rating.IsPublic = req.IsPublic

	return s.repo.UpdateRating(rating)
}

// DeleteRating 删除评价
func (s *RatingService) DeleteRating(id uint) error {
	return s.repo.DeleteRating(id)
}
