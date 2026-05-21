package service

import (
	"context"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/model"
	"agussyahrilmubarok.github.io/backend/internal/repository"
	"agussyahrilmubarok.github.io/backend/pkg/helper"
)

//go:generate mockery --name=IAuthService
type IAuthService interface {
	SignUp(ctx context.Context, param model.SignUpRequest) (*model.AuthResponse, error)
	SignIn(ctx context.Context, param model.SignInRequest) (*model.AuthResponse, error)
}

type authService struct {
	userRepository repository.IUserRepository
}

// SignUp implements [IAuthService].
func (s *authService) SignUp(ctx context.Context, param model.SignUpRequest) (*model.AuthResponse, error) {
	exist, err := s.userRepository.FindByEmail(ctx, param.Email)
	if exist != nil || err != nil {
		return nil, domain.ErrUserEmailExists
	}

	hashPassword, err := helper.PasswordHash(param.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: hashPassword,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: "token",
		User:  model.ToUserResponse(user),
	}, nil
}

// SignIn implements [IAuthService].
func (s *authService) SignIn(ctx context.Context, param model.SignInRequest) (*model.AuthResponse, error) {
	panic("unimplemented")
}

func NewAuthService(
	userRepository repository.IUserRepository,
) IAuthService {
	return &authService{
		userRepository: userRepository,
	}
}
