package repository

import (
	"dididaren/internal/model"
	"time"

	"gorm.io/gorm"
)

type SecurityRepository struct {
	db *gorm.DB
}

func NewSecurityRepository(db *gorm.DB) *SecurityRepository {
	return &SecurityRepository{
		db: db,
	}
}

// CreateStaff 创建安保人员
func (r *SecurityRepository) CreateStaff(staff *model.Staff) error {
	return r.db.Create(staff).Error
}

// GetStaffByID 根据ID获取安保人员
func (r *SecurityRepository) GetStaffByID(id uint) (*model.Staff, error) {
	var staff model.Staff
	err := r.db.First(&staff, id).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

// GetStaffByUserID 根据用户ID获取安保人员
func (r *SecurityRepository) GetStaffByUserID(userID uint) (*model.Staff, error) {
	var staff model.Staff
	err := r.db.Where("user_id = ?", userID).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

// ListStaffs 获取安保人员列表
func (r *SecurityRepository) ListStaffs(page, size int, status string) ([]model.Staff, int64, error) {
	var staffs []model.Staff
	var total int64

	query := r.db.Model(&model.Staff{})
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * size).Limit(size).Find(&staffs).Error
	if err != nil {
		return nil, 0, err
	}

	return staffs, total, nil
}

// UpdateStaffStatus 更新安保人员状态
func (r *SecurityRepository) UpdateStaffStatus(id uint, status string) error {
	return r.db.Model(&model.Staff{}).Where("id = ?", id).Update("status", status).Error
}

// CreateRating 创建评价
func (r *SecurityRepository) CreateRating(rating *model.Rating) error {
	return r.db.Create(rating).Error
}

// ListRatings 获取评价列表
func (r *SecurityRepository) ListRatings(staffID uint) ([]model.Rating, error) {
	var ratings []model.Rating
	err := r.db.Where("staff_id = ?", staffID).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

// UpdateLocation 更新安保人员位置
func (r *SecurityRepository) UpdateLocation(staffID uint, lat, lng float64) error {
	return r.db.Model(&model.Staff{}).Where("id = ?", staffID).Updates(map[string]interface{}{
		"latitude":    lat,
		"longitude":   lng,
		"last_active": time.Now(),
	}).Error
}

// Create 创建安保人员
func (r *SecurityRepository) Create(staff *model.SecurityStaff) error {
	return r.db.Create(staff).Error
}

// GetByID 根据ID获取安保人员
func (r *SecurityRepository) GetByID(id uint) (*model.SecurityStaff, error) {
	var staff model.SecurityStaff
	if err := r.db.First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

// GetByUserID 根据用户ID获取安保人员
func (r *SecurityRepository) GetByUserID(userID uint) (*model.SecurityStaff, error) {
	var staff model.SecurityStaff
	if err := r.db.Where("user_id = ?", userID).First(&staff).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}

// Update 更新安保人员
func (r *SecurityRepository) Update(staff *model.SecurityStaff) error {
	return r.db.Save(staff).Error
}

// GetEventByID 根据ID获取事件
func (r *SecurityRepository) GetEventByID(id uint) (*model.Emergency, error) {
	var event model.Emergency
	if err := r.db.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

// UpdateEvent 更新事件
func (r *SecurityRepository) UpdateEvent(event *model.Emergency) error {
	return r.db.Save(event).Error
}

// CreateHandlingRecord 创建处理记录
func (r *SecurityRepository) CreateHandlingRecord(record *model.HandlingRecord) error {
	return r.db.Create(record).Error
}

// UpdateOnlineStatus 更新安保人员在线状态
func (r *SecurityRepository) UpdateOnlineStatus(userID uint, isOnline bool) error {
	return r.db.Model(&model.SecurityStaff{}).
		Where("user_id = ?", userID).
		Update("is_online", isOnline).Error
}

// GetNearbyStaff 获取附近的安保人员
func (r *SecurityRepository) GetNearbyStaff(latitude, longitude float64, radius float64) ([]*model.SecurityStaff, error) {
	var staff []*model.SecurityStaff
	// 使用简单的经纬度范围查询，实际项目中可能需要使用更复杂的空间查询
	err := r.db.Where("is_online = ? AND latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?",
		true,
		latitude-radius, latitude+radius,
		longitude-radius, longitude+radius,
	).Find(&staff).Error
	if err != nil {
		return nil, err
	}
	return staff, nil
}

// IncrementOrderCount 增加安保人员接单数
func (r *SecurityRepository) IncrementOrderCount(userID uint) error {
	return r.db.Model(&model.SecurityStaff{}).
		Where("user_id = ?", userID).
		UpdateColumn("order_count", gorm.Expr("order_count + ?", 1)).Error
}
