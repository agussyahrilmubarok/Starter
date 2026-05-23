package service

import (
	"context"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/model"
	"agussyahrilmubarok.github.io/backend/internal/repository"
	"agussyahrilmubarok.github.io/backend/pkg/helper"
	"agussyahrilmubarok.github.io/backend/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
	log := logger.FromCtx(ctx)

	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		log.Error("user: failed to get all users", zap.Error(err))
		return nil, err
	}

	log.Debug("user: get all users", zap.Int("count", len(users)))

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = model.ToUserResponse(&user)
	}
	return responses, nil
}

// GetByID implements [IUserService].
func (s *userService) GetByID(ctx context.Context, ID string) (*model.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_id", ID))

	user, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		log.Error("user: failed to find user by id", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("user: user not found")
		return nil, domain.ErrUserNotFound
	}

	log.Debug("user: get by id success")
	res := model.ToUserResponse(user)
	return &res, nil
}

// Create implements [IUserService].
func (s *userService) Create(ctx context.Context, param model.CreateUserRequest) (*model.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_email", param.Email))

	exist, err := s.userRepository.FindByEmail(ctx, param.Email)
	if err != nil {
		log.Error("user: failed to check existing email", zap.Error(err))
		return nil, err
	}
	if exist != nil {
		log.Warn("user: email already exists")
		return nil, domain.ErrUserEmailExists
	}

	hashPassword, err := helper.PasswordHash(param.Password)
	if err != nil {
		log.Error("user: failed to hash password", zap.Error(err))
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.New().String(),
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: hashPassword,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		log.Error("user: failed to create user", zap.Error(err))
		return nil, err
	}

	log.Info("user: created successfully", zap.String("user_id", user.ID))
	res := model.ToUserResponse(user)
	return &res, nil
}

// UpdateByID implements [IUserService].
func (s *userService) UpdateByID(ctx context.Context, ID string, param model.UpdateUserRequest) (*model.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_id", ID))

	user, err := s.userRepository.FindByID(ctx, ID)
	if err != nil {
		log.Error("user: failed to find user for update", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("user: user not found for update")
		return nil, domain.ErrUserNotFound
	}

	if param.Name != "" {
		user.Name = param.Name
	}
	if param.Email != "" && param.Email != user.Email {
		exist, _ := s.userRepository.FindByEmail(ctx, param.Email)
		if err != nil {
			return nil, err
		}
		if exist != nil {
			return nil, domain.ErrUserEmailExists
		}
		user.Email = strings.ToLower(param.Email)
	}
	if param.Password != "" {
		hashPassword, err := helper.PasswordHash(param.Password)
		if err != nil {
			log.Error("user: failed to hash password for update", zap.Error(err))
			return nil, err
		}
		user.Password = hashPassword
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		log.Error("user: failed to save updated user", zap.Error(err))
		return nil, err
	}

	log.Info("user: updated successfully")
	res := model.ToUserResponse(user)
	return &res, nil
}

// DeleteByID implements [IUserService].
func (s *userService) DeleteByID(ctx context.Context, ID string) error {
	log := logger.FromCtx(ctx).With(zap.String("user_id", ID))

	if err := s.userRepository.DeleteByID(ctx, ID); err != nil {
		log.Error("user: failed to delete user", zap.Error(err))
		return err
	}

	log.Info("user: deleted successfully")
	return nil
}

func NewUserService(
	userRepository repository.IUserRepository,
) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}
