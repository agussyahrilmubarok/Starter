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

//go:generate mockery --name=IAuthService
type IAuthService interface {
	SignUp(ctx context.Context, param model.SignUpRequest) (*model.AuthResponse, error)
	SignIn(ctx context.Context, param model.SignInRequest) (*model.AuthResponse, error)
}

type authService struct {
	userRepository repository.IUserRepository
	jwtService     IJWTService
}

// SignUp implements [IAuthService].
func (s *authService) SignUp(ctx context.Context, param model.SignUpRequest) (*model.AuthResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_email", param.Email))

	exist, err := s.userRepository.FindByEmail(ctx, param.Email)
	if err != nil {
		log.Error("sign-up: failed to find user by email", zap.Error(err))
		return nil, err
	}
	if exist != nil {
		log.Warn("sign-up: email already exists")
		return nil, domain.ErrUserEmailExists
	}

	hashPassword, err := helper.PasswordHash(param.Password)
	if err != nil {
		log.Error("sign-up: failed to hash password", zap.Error(err))
		return nil, err
	}

	user := &domain.User{
		ID:       uuid.New().String(),
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: hashPassword,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		log.Error("sign-up: failed to create user", zap.Error(err))
		return nil, err
	}

	tokenString, err := s.jwtService.Generate(ctx, user.ID)
	if err != nil {
		log.Error("sign-up: failed to generate token, rolling back user", zap.String("user_id", user.ID), zap.Error(err))
		if err := s.userRepository.DeleteByID(ctx, user.ID); err != nil {
			log.Error("sign-up: failed to rollback user creation", zap.Error(err))
		}
		return nil, err
	}

	log.Info("sign-up: user created successfully", zap.String("user_id", user.ID))
	return &model.AuthResponse{
		Token: tokenString,
		User:  model.ToUserResponse(user),
	}, nil
}

// SignIn implements [IAuthService].
func (s *authService) SignIn(ctx context.Context, param model.SignInRequest) (*model.AuthResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("user_email", param.Email))

	user, err := s.userRepository.FindByEmail(ctx, strings.ToLower(param.Email))
	if err != nil {
		log.Error("sign-in: failed to find user by email", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("sign-in: email not found")
		return nil, domain.ErrUserEmailNotFound
	}

	if ok := helper.PasswordCheck(param.Password, user.Password); !ok {
		log.Warn("sign-in: password not match", zap.String("user_id", user.ID))
		return nil, domain.ErrUserPasswordNotMatch
	}

	tokenString, err := s.jwtService.Generate(ctx, user.ID)
	if err != nil {
		log.Error("sign-in: failed to generate token", zap.String("user_id", user.ID), zap.Error(err))
		return nil, err
	}

	log.Info("sign-in: success", zap.String("user_id", user.ID))
	return &model.AuthResponse{
		Token: tokenString,
		User:  model.ToUserResponse(user),
	}, nil
}

func NewAuthService(
	userRepository repository.IUserRepository,
	jwtService IJWTService,
) IAuthService {
	return &authService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}
