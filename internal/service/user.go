package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Register 用户注册
func (s *UserService) Register(req *model.RegisterRequest) (*model.User, error) {
	// 检查手机号是否已注册
	if s.repo.ExistsByPhone(req.Phone) {
		return nil, errors.New("该手机号已注册")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Phone:     req.Phone,
		Password:  string(hashedPassword),
		Name:      req.Name,
		Status:    1,
		LastLogin: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(req *model.LoginRequest) (*model.User, error) {
	user, err := s.repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}

	// 更新最后登录时间
	if err := s.repo.UpdateLastLogin(user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

// GetProfile 获取用户信息
func (s *UserService) GetProfile(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(id uint, req *model.UpdateProfileRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
