package service

import (
	"context"
	"strings"

	"agussyahrilmubarok.github.io/web/internal/domain"
	"agussyahrilmubarok.github.io/web/internal/model"
	"agussyahrilmubarok.github.io/web/internal/repository"
	"agussyahrilmubarok.github.io/web/pkg/cryptoutil"
	"agussyahrilmubarok.github.io/web/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

//go:generate mockery --name=IAuthService
type IAuthService interface {
	SignUp(ctx context.Context, param model.SignUpRequest) (*model.UserResponse, error)
	SignIn(ctx context.Context, param model.SignInRequest) (*model.UserResponse, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

// SignUp implements [IAuthService].
func (s *authService) SignUp(ctx context.Context, param model.SignUpRequest) (*model.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_email", param.Email))

	exist, err := s.userRepository.ExistsByEmailIgnorecase(ctx, param.Email)
	if err != nil {
		log.Error("failed to check existing user email", zap.Error(err))
		return nil, err
	}

	if exist {
		log.Warn("user email already exists")
		return nil, domain.ErrUserEmailExists
	}

	hashPassword, err := cryptoutil.PasswordHash(param.Password)
	if err != nil {
		log.Error("failed to hash user password", zap.Error(err))
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.New(),
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: hashPassword,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		log.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	log.Info("user signup completed successfully", zap.String("user_id", user.ID.String()))
	userResponse := model.ToUserResponse(user)
	return &userResponse, nil
}

// SignIn implements [IAuthService].
func (s *authService) SignIn(ctx context.Context, param model.SignInRequest) (*model.UserResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_email", param.Email))

	user, err := s.userRepository.FindByEmail(ctx, strings.ToLower(param.Email))
	if err != nil {
		log.Error("failed to find user by email", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("user email not found")
		return nil, domain.ErrUserEmailNotFound
	}

	if err := cryptoutil.PasswordCheck(param.Password, user.Password); err != nil {
		log.Warn("invalid user password", zap.String("user_id", user.ID.String()))
		return nil, domain.ErrUserPasswordNotMatch
	}

	log.Info("user signin completed successfully", zap.String("user_id", user.ID.String()))
	userResponse := model.ToUserResponse(user)
	return &userResponse, nil
}

func NewAuthService(
	userRepository repository.IUserRepository,
) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}
