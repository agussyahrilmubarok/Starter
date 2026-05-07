package usecase

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/pkg/helper"
	"context"
	"fmt"
	"strings"
)

//go:generate mockery --name=IAuthUseCase
type IAuthUseCase interface {
	Register(ctx context.Context, param *dto.UserCreateRequest) (*dto.UserDTO, error)
	Login(ctx context.Context, param *dto.UserLoginRequest) (*dto.AuthDTO, error)
}

type authUseCase struct {
	userRepository repository.IUserRepository
}

// Register implements [IAuthUseCase].
func (uc *authUseCase) Register(ctx context.Context, param *dto.UserCreateRequest) (*dto.UserDTO, error) {
	user := domain.User{
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: helper.PasswordEncrypt(param.Password),
	}

	if err := uc.userRepository.Create(ctx, &user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "email") {
				return nil, domain.ErrUserEmailUsed
			}
		}
		return nil, err
	}

	var userDto dto.UserDTO
	userDto.FromDomain(&user)

	return &userDto, nil
}

// Login implements [IAuthUseCase].
func (uc *authUseCase) Login(ctx context.Context, param *dto.UserLoginRequest) (*dto.AuthDTO, error) {
	user, err := uc.userRepository.FindByEmail(ctx, param.Email)
	if user == nil && err == nil {
		return nil, domain.ErrUserNotFound
	}

	if err := helper.PasswordCompare(user.Password, param.Password); err != nil {
		return nil, domain.ErrUserPasswordInvalid
	}

	token := helper.GenerateToken(fmt.Sprint(user.ID))
	var userDto dto.UserDTO
	userDto.FromDomain(user)

	return &dto.AuthDTO{
		Token: token,
		User:  userDto,
	}, nil
}

func NewAuthUseCase(
	userRepository repository.IUserRepository,
) IAuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
	}
}
