package service

import (
	"dididaren/internal/model"
	"dididaren/internal/repository"
	"dididaren/pkg/auth"
	"dididaren/pkg/errors"
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
	exists := s.repo.ExistsByPhone(req.Phone)
	if exists {
		return nil, errors.ErrPhoneAlreadyRegistered
	}

	// 创建用户
	user := &model.User{
		Phone:    req.Phone,
		Password: auth.HashPassword(req.Password),
		Name:     req.Name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(req *model.LoginRequest) (string, error) {
	// 获取用户信息
	user, err := s.repo.GetByPhone(req.Phone)
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	// 验证密码
	if !auth.VerifyPassword(req.Password, user.Password) {
		return "", errors.ErrInvalidCredentials
	}

	// 生成token
	token, err := auth.GenerateToken(user.ID, user.Phone, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id uint, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	user.Avatar = req.Avatar

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePassword 更新密码
func (s *UserService) UpdatePassword(id uint, req *model.UpdatePasswordRequest) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// 验证旧密码
	if !auth.VerifyPassword(req.OldPassword, user.Password) {
		return errors.ErrInvalidPassword
	}

	// 更新密码
	user.Password = auth.HashPassword(req.NewPassword)
	return s.repo.Update(user)
}

func (s *UserService) List(page, size int) ([]model.User, int64, error) {
	return s.repo.List(page, size)
}

func (s *UserService) Delete(userID uint) error {
	return s.repo.Delete(userID)
}
