package usecase

import (
	"context"
	"errors"
	"strings"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"agussyahrilmubarok.github.io/backend/pkg/crypto"
	"agussyahrilmubarok.github.io/backend/pkg/logger"
	"go.uber.org/zap"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)

//go:generate mockery --name=AuthUseCase
type AuthUseCase interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.AuthResponse, error)
	SignIn(ctx context.Context, req dto.SignInRequest) (*dto.AuthResponse, error)
}

type authUseCase struct {
	userRepository domain.UserRepository
	jwtManager     security.JWTManager
}

func NewAuthUseCase(userRepository domain.UserRepository, jwtManager security.JWTManager) AuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
		jwtManager:     jwtManager,
	}
}

// SignUp implements [AuthUseCase].
func (uc *authUseCase) SignUp(ctx context.Context, req dto.SignUpRequest) (*dto.AuthResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("email", req.Email))

	exists, err := uc.userRepository.ExistsByEmail(ctx, req.Email)
	if err != nil {
		log.Error("failed to check email existence", zap.Error(err))
		return nil, err
	}
	if exists {
		log.Warn("email already exists")
		return nil, ErrEmailAlreadyExist
	}

	hashed, err := crypto.HashPassword(req.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		return nil, err
	}

	user := domain.User{
		Name:     req.Name,
		Email:    strings.ToLower(req.Email),
		Password: hashed,
	}

	if err := uc.userRepository.Create(ctx, &user); err != nil {
		log.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	token, err := uc.jwtManager.GenerateToken(user.ID)
	if err != nil {
		log.Error("failed to generate token", zap.Error(err))
		return nil, err
	}

	log.Info("user signed up", zap.String("id", user.ID.String()))
	return &dto.AuthResponse{
		Token: token,
		User:  dto.NewUserResponse(&user),
	}, nil
}

// SignIn implements [AuthUseCase].
func (uc *authUseCase) SignIn(ctx context.Context, req dto.SignInRequest) (*dto.AuthResponse, error) {
	log := logger.FromCtx(ctx).With(zap.String("email", req.Email))

	user, err := uc.userRepository.FindByEmail(ctx, strings.ToLower(req.Email))
	if err != nil {
		log.Error("failed to find user", zap.Error(err))
		return nil, err
	}
	if user == nil {
		log.Warn("invalid email")
		return nil, ErrInvalidEmail
	}

	if !crypto.CheckPassword(user.Password, req.Password) {
		log.Warn("invalid password")
		return nil, ErrInvalidPassword
	}

	token, err := uc.jwtManager.GenerateToken(user.ID)
	if err != nil {
		log.Error("failed to generate token", zap.Error(err))
		return nil, err
	}

	log.Info("user signed in", zap.String("id", user.ID.String()))
	return &dto.AuthResponse{
		Token: token,
		User:  dto.NewUserResponse(user),
	}, nil
}
