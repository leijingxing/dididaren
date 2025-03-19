package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"dididaren/pkg/errors"
)

type SecurityService struct {
	staffRepo *repository.SecurityStaffRepository
	eventRepo *repository.EmergencyEventRepository
	userRepo  *repository.UserRepository
}

func NewSecurityService(
	staffRepo *repository.SecurityStaffRepository,
	eventRepo *repository.EmergencyEventRepository,
	userRepo *repository.UserRepository,
) *SecurityService {
	return &SecurityService{
		staffRepo: staffRepo,
		eventRepo: eventRepo,
		userRepo:  userRepo,
	}
}

// ApplySecurityStaff 申请成为安保人员
func (s *SecurityService) ApplySecurityStaff(userID uint, companyName, licenseNumber string, certFiles []string) (*model.SecurityStaff, error) {
	// 检查用户是否存在
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 检查是否已经是安保人员
	staff, err := s.staffRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if staff != nil {
		return nil, errors.New("已经是安保人员")
	}

	// 创建安保人员记录
	staff = &model.SecurityStaff{
		UserID:              userID,
		CompanyName:         companyName,
		LicenseNumber:       licenseNumber,
		CertificationStatus: 0, // 待审核
		Rating:              5.0,
		TotalOrders:         0,
		OnlineStatus:        0,
	}

	if err := s.staffRepo.Create(staff); err != nil {
		return nil, err
	}

	return staff, nil
}

// UpdateLocation 更新位置信息
func (s *SecurityService) UpdateLocation(userID uint, locationLat, locationLng float64, onlineStatus int8) error {
	staff, err := s.staffRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if staff == nil {
		return errors.New("不是安保人员")
	}

	staff.CurrentLocationLat = locationLat
	staff.CurrentLocationLng = locationLng
	staff.OnlineStatus = onlineStatus

	if err := s.staffRepo.Update(staff); err != nil {
		return err
	}

	return nil
}

// AcceptEvent 接受事件
func (s *SecurityService) AcceptEvent(userID uint, eventID string) error {
	// 检查安保人员状态
	staff, err := s.staffRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if staff == nil {
		return errors.New("不是安保人员")
	}
	if staff.OnlineStatus != 1 {
		return errors.New("安保人员不在线")
	}

	// 获取事件信息
	event, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		return err
	}
	if event == nil {
		return errors.New("事件不存在")
	}
	if event.Status != 0 {
		return errors.New("事件状态不允许接单")
	}

	// 更新事件状态
	event.Status = 1 // 处理中
	if err := s.eventRepo.Update(event); err != nil {
		return err
	}

	// 创建处理记录
	record := &model.EventHandlingRecord{
		EventID:     event.ID,
		HandlerID:   userID,
		HandlerType: 2, // 安保人员
		ActionType:  1, // 接单
	}

	if err := s.eventRepo.CreateHandlingRecord(record); err != nil {
		return err
	}

	return nil
}

// CompleteEvent 完成事件
func (s *SecurityService) CompleteEvent(userID uint, eventID string, remark string) error {
	// 检查安保人员状态
	staff, err := s.staffRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if staff == nil {
		return errors.New("不是安保人员")
	}

	// 获取事件信息
	event, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		return err
	}
	if event == nil {
		return errors.New("事件不存在")
	}
	if event.Status != 1 {
		return errors.New("事件状态不允许完成")
	}

	// 更新事件状态
	event.Status = 2 // 已完成
	if err := s.eventRepo.Update(event); err != nil {
		return err
	}

	// 创建处理记录
	record := &model.EventHandlingRecord{
		EventID:     event.ID,
		HandlerID:   userID,
		HandlerType: 2, // 安保人员
		ActionType:  4, // 完成
		Remark:      remark,
	}

	if err := s.eventRepo.CreateHandlingRecord(record); err != nil {
		return err
	}

	// 更新安保人员订单数
	staff.TotalOrders++
	if err := s.staffRepo.Update(staff); err != nil {
		return err
	}

	return nil
}

// GetStaffInfo 获取安保人员信息
func (s *SecurityService) GetStaffInfo(userID uint) (*model.SecurityStaff, error) {
	staff, err := s.staffRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if staff == nil {
		return nil, errors.New("不是安保人员")
	}

	return staff, nil
}
