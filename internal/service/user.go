package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Register 用户注册
func (s *UserService) Register(req *model.RegisterRequest) (*model.User, error) {
	// 检查手机号是否已注册
	existingUser, err := s.repo.GetByPhone(req.Phone)
	if err == nil && existingUser != nil {
		return nil, errors.New("手机号已注册")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: string(hashedPassword),
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
func (s *UserService) GetProfile(userID uint) (*model.User, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(userID uint, req *model.UpdateProfileRequest) (*model.User, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(page, size int) ([]model.User, int64, error) {
	return s.repo.List(page, size)
}

func (s *UserService) Delete(userID uint) error {
	return s.repo.Delete(userID)
}
