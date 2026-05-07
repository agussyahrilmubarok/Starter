package usecase

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/repository"
	"backend/pkg/exception"
	"backend/pkg/helper"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

//go:generate mockery --name=IUserUseCase
type IUserUseCase interface {
	FindAll(ctx context.Context) ([]dto.UserDTO, error)
	FindByID(ctx context.Context, userID uint) (*dto.UserDTO, error)
	Create(ctx context.Context, param *dto.UserCreateRequest) (*dto.UserDTO, error)
	UpdateByID(ctx context.Context, userID uint, param *dto.UserUpdateRequest) (*dto.UserDTO, error)
	DeleteByID(ctx context.Context, userID uint) error
}

type userUseCase struct {
	userRepository repository.IUserRepository
}

// FindAll implements [IUserUseCase].
func (uc *userUseCase) FindAll(ctx context.Context) ([]dto.UserDTO, error) {
	var userDtos []dto.UserDTO

	users, err := uc.userRepository.FindAll(ctx)
	if err != nil {
		return nil, exception.NewBadRequest("User not found", exception.HttpErrMap("Error", "No records"))
	}

	for _, user := range users {
		var userDto dto.UserDTO
		userDto.FromDomain(&user)
		userDtos = append(userDtos, userDto)
	}

	return userDtos, nil
}

// FindByID implements [IUserUseCase].
func (uc *userUseCase) FindByID(ctx context.Context, userID uint) (*dto.UserDTO, error) {
	user, err := uc.userRepository.FindByID(ctx, userID)
	if user == nil && err == nil {
		return nil, exception.NewBadRequest("User not found", exception.HttpErrMap("Error", "User not found"))
	}

	var userDto dto.UserDTO
	userDto.FromDomain(user)

	return &userDto, nil
}

// Create implements [IUserUseCase].
func (uc *userUseCase) Create(ctx context.Context, param *dto.UserCreateRequest) (*dto.UserDTO, error) {
	user := domain.User{
		Name:     param.Name,
		Email:    strings.ToLower(param.Email),
		Password: helper.PasswordEncrypt(param.Password),
	}

	if err := uc.userRepository.Create(ctx, &user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "email") {
				return nil, exception.NewUnprocessableEntity("Email already used", exception.HttpErrMap("Email", "Email already exists"))
			}
		}
		return nil, err
	}

	var userDto dto.UserDTO
	userDto.FromDomain(&user)

	return &userDto, nil
}

// UpdateByID implements [IUserUseCase].
func (uc *userUseCase) UpdateByID(ctx context.Context, userID uint, param *dto.UserUpdateRequest) (*dto.UserDTO, error) {
	user, err := uc.userRepository.FindByID(ctx, userID)
	if user == nil && err == nil {
		return nil, exception.NewBadRequest("User not found", exception.HttpErrMap("Error", "User not found"))
	}

	user.Name = param.Name
	user.Email = param.Email
	if param.Password != "" {
		user.Password = helper.PasswordEncrypt(param.Password)
	}

	if err := uc.userRepository.Save(ctx, user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "email") {
				return nil, exception.NewConflict("Email already used", exception.HttpErrMap("Email", "Email already exists"))
			}
		}
		return nil, err
	}

	var userDto dto.UserDTO
	userDto.FromDomain(user)

	return &userDto, nil
}

// DeleteByID implements [IUserUseCase].
func (uc *userUseCase) DeleteByID(ctx context.Context, userID uint) error {
	if err := uc.userRepository.DeleteByID(ctx, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.NewBadRequest("User not found", exception.HttpErrMap("Error", "Record is not found"))
		}
		return err
	}

	return nil
}

func NewUserUseCase(
	userRepository repository.IUserRepository,
) IUserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}
