package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

// CreateRating 创建评价
func (r *RatingRepository) CreateRating(rating *model.Rating) (*model.Rating, error) {
	err := r.db.Create(rating).Error
	if err != nil {
		return nil, err
	}
	return rating, nil
}

// GetRatingByID 根据ID获取评价
func (r *RatingRepository) GetRatingByID(id uint) (*model.Rating, error) {
	var rating model.Rating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

// GetByEventID 获取事件相关的评价
func (r *RatingRepository) GetByEventID(eventID uint) (*model.Rating, error) {
	var rating model.Rating
	err := r.db.Where("event_id = ?", eventID).First(&rating).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &rating, nil
}

// ListRatings 获取安保人员的评价列表
func (r *RatingRepository) ListRatings(staffID uint) ([]*model.Rating, error) {
	var ratings []*model.Rating
	err := r.db.Where("staff_id = ?", staffID).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

// CalculateStaffAverageRating 计算安保人员的平均评分
func (r *RatingRepository) CalculateStaffAverageRating(staffID uint) (float64, error) {
	var avg float64
	err := r.db.Model(&model.Rating{}).
		Where("staff_id = ?", staffID).
		Select("AVG(score)").
		Scan(&avg).Error
	if err != nil {
		return 0, err
	}
	return avg, nil
}

// UpdateRating 更新评价
func (r *RatingRepository) UpdateRating(rating *model.Rating) error {
	return r.db.Save(rating).Error
}

// DeleteRating 删除评价
func (r *RatingRepository) DeleteRating(id uint) error {
	return r.db.Delete(&model.Rating{}, id).Error
}
