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

	tokenString, err := s.jwtService.Generate(ctx, user.ID)
	if err != nil {
		s.userRepository.DeleteByID(ctx, user.ID)
		return nil, err
	}

	return &model.AuthResponse{
		Token: tokenString,
		User:  model.ToUserResponse(user),
	}, nil
}

// SignIn implements [IAuthService].
func (s *authService) SignIn(ctx context.Context, param model.SignInRequest) (*model.AuthResponse, error) {
	user, err := s.userRepository.FindByEmail(ctx, strings.ToLower(param.Email))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserEmailNotFound
	}

	if ok := helper.PasswordCheck(param.Password, user.Password); !ok {
		return nil, domain.ErrUserPasswordNotMatch
	}

	tokenString, err := s.jwtService.Generate(ctx, user.ID)
	if err != nil {
		return nil, err
	}

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
