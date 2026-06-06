package usecase

import (
	"context"
	"errors"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/pkg/crypto"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyInUse = errors.New("email is already in use")
)

//go:generate mockery --name=UserUseCase
type UserUseCase interface {
	GetAll(ctx context.Context) ([]dto.UserResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error)
	Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

// GetAll implements [UserUseCase].
func (uc *userUseCase) GetAll(ctx context.Context) ([]dto.UserResponse, error) {
	log := logger.FromCtx(ctx)

	users, err := uc.userRepository.FindAll(ctx)
	if err != nil {
		log.Error("failed to get all users", zap.Error(err))
		return nil, err
	}

	log.Info("users fetched", zap.Int("count", len(users)))
	return dto.NewUserListResponse(users), nil
}

// GetByID implements [UserUseCase].
func (uc *userUseCase) GetByID(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("id", id.String()))

	user, err := uc.userRepository.FindByID(ctx, id)
	if err != nil {
		log.Error("failed to get user", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("user not found")
		return nil, ErrUserNotFound
	}

	log.Info("user fetched", zap.String("user_id", user.ID.String()))
	res := dto.NewUserResponse(user)
	return &res, nil
}

// Create implements [UserUseCase].
func (uc *userUseCase) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("email", req.Email))

	exists, err := uc.userRepository.ExistsByEmail(ctx, strings.ToLower(req.Email))
	if err != nil {
		log.Error("failed to check email existence", zap.Error(err))
		return nil, err
	}
	if exists {
		log.Warn("email already exists")
		return nil, ErrEmailAlreadyInUse
	}

	hashed, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		return nil, err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    strings.ToLower(req.Email),
		Password: hashed,
	}

	if err := uc.userRepository.Create(ctx, user); err != nil {
		log.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	log.Info("user created", zap.String("id", user.ID.String()))
	res := dto.NewUserResponse(user)
	return &res, nil
}

// Update implements [UserUseCase].
func (uc *userUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("id", id.String()))

	user, err := uc.userRepository.FindByID(ctx, id)
	if err != nil {
		log.Error("failed to get user", zap.String("user_id", id.String()), zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("user not found", zap.String("user_id", id.String()))
		return nil, ErrUserNotFound
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" && req.Email != user.Email {
		exists, err := uc.userRepository.ExistsByEmail(ctx, strings.ToLower(req.Email))
		if err != nil {
			log.Error("failed to check email existence", zap.Error(err))
			return nil, err
		}
		if exists {
			log.Warn("email already exists", zap.String("email", req.Email))
			return nil, ErrEmailAlreadyInUse
		}
		user.Email = strings.ToLower(req.Email)
	}

	if req.Password != "" {
		hashed, err := crypto.HashPassword(req.Password)
		if err != nil {
			log.Error("failed to hash password", zap.Error(err))
			return nil, err
		}
		user.Password = hashed
	}

	if err := uc.userRepository.Update(ctx, user); err != nil {
		log.Error("failed to update user", zap.Error(err))
		return nil, err
	}

	log.Info("user updated", zap.String("user_id", user.ID.String()))
	res := dto.NewUserResponse(user)
	return &res, nil
}

// Delete implements [UserUseCase].
func (uc *userUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	log := logger.FromCtx(ctx).With(zap.String("id", id.String()))

	user, err := uc.userRepository.FindByID(ctx, id)
	if err != nil {
		log.Error("failed to get user", zap.Error(err))
		return err
	}
	if user == nil {
		log.Warn("user not found", zap.String("user_id", id.String()))
		return ErrUserNotFound
	}

	if err := uc.userRepository.Delete(ctx, id); err != nil {
		log.Error("failed to delete user", zap.Error(err))
		return err
	}

	log.Info("user deleted", zap.String("user_id", user.ID.String()))
	return nil
}
