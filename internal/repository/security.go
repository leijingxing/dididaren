package repository

import (
	"dididaren/internal/model"

	"gorm.io/gorm"
)

type SecurityStaffRepository struct {
	db *gorm.DB
}

func NewSecurityStaffRepository(db *gorm.DB) *SecurityStaffRepository {
	return &SecurityStaffRepository{
		db: db,
	}
}

// Create 创建安保人员记录
func (r *SecurityStaffRepository) Create(staff *model.SecurityStaff) error {
	return r.db.Create(staff).Error
}

// GetByUserID 根据用户ID获取安保人员信息
func (r *SecurityStaffRepository) GetByUserID(userID uint) (*model.SecurityStaff, error) {
	var staff model.SecurityStaff
	err := r.db.Where("user_id = ?", userID).First(&staff).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &staff, nil
}

// Update 更新安保人员信息
func (r *SecurityStaffRepository) Update(staff *model.SecurityStaff) error {
	return r.db.Save(staff).Error
}

// UpdateLocation 更新安保人员位置
func (r *SecurityStaffRepository) UpdateLocation(userID uint, latitude, longitude float64) error {
	return r.db.Model(&model.SecurityStaff{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"latitude":  latitude,
			"longitude": longitude,
		}).Error
}

// UpdateOnlineStatus 更新安保人员在线状态
func (r *SecurityStaffRepository) UpdateOnlineStatus(userID uint, isOnline bool) error {
	return r.db.Model(&model.SecurityStaff{}).
		Where("user_id = ?", userID).
		Update("is_online", isOnline).Error
}

// GetNearbyStaff 获取附近的安保人员
func (r *SecurityStaffRepository) GetNearbyStaff(latitude, longitude float64, radius float64) ([]*model.SecurityStaff, error) {
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
func (r *SecurityStaffRepository) IncrementOrderCount(userID uint) error {
	return r.db.Model(&model.SecurityStaff{}).
		Where("user_id = ?", userID).
		UpdateColumn("order_count", gorm.Expr("order_count + ?", 1)).Error
}
