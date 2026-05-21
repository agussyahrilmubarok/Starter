package service

import (
	"context"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/model"
	"agussyahrilmubarok.github.io/backend/internal/repository"
	"agussyahrilmubarok.github.io/backend/pkg/helper"
	"github.com/google/uuid"
)

//go:generate mockery --name=IUserService
type IUserService interface {
	GetAll(ctx context.Context) ([]model.UserResponse, error)
	GetByID(ctx context.Context, ID string) (*model.UserResponse, error)
	Create(ctx context.Context, param model.CreateUserRequest) (*model.UserResponse, error)
	UpdateByID(ctx context.Context, ID string, param model.UpdateUserRequest) (*model.UserResponse, error)
	DeleteByID(ctx context.Context, ID string) error
}

type userService struct {
	userRepository repository.IUserRepository
}

// GetAll implements [IUserService].
func (s *userService) GetAll(ctx context.Context) ([]model.UserResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = model.ToUserResponse(&user)
	}

	return responses, nil
}

// GetByID implements [IUserService].
func (s *userService) GetByID(ctx context.Context, ID string) (*model.UserResponse, error) {
	user, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	res := model.ToUserResponse(user)
	return &res, nil
}

// Create implements [IUserService].
func (s *userService) Create(ctx context.Context, param model.CreateUserRequest) (*model.UserResponse, error) {
	exist, err := s.userRepository.FindByEmail(ctx, param.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, domain.ErrUserEmailExists
	}

	hashPassword, err := helper.PasswordHash(param.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.New().String(),
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: hashPassword,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	res := model.ToUserResponse(user)
	return &res, nil
}

// UpdateByID implements [IUserService].
func (s *userService) UpdateByID(ctx context.Context, ID string, param model.UpdateUserRequest) (*model.UserResponse, error) {
	user, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	if param.Name != "" {
		user.Name = param.Name
	}

	if param.Email != "" {
		user.Email = param.Email
	}

	if param.Password != "" {
		hashPassword, err := helper.PasswordHash(param.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashPassword
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		return nil, err
	}

	res := model.ToUserResponse(user)
	return &res, nil
}

// DeleteByID implements [IUserService].
func (s *userService) DeleteByID(ctx context.Context, ID string) error {
	user, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrUserNotFound
	}

	return s.userRepository.DeleteByID(ctx, ID)
}

func NewUserService(
	userRepository repository.IUserRepository,
) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}
