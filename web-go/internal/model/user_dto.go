package model

import (
	"time"

	"agussyahrilmubarok.github.io/web/internal/domain"
)

type UserResponse struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type CreateUserRequest struct {
	Name     string `form:"name"     validate:"required,min=2,max=100"`
	Email    string `form:"email"    validate:"required,email,max=100"`
	Password string `form:"password" validate:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	Name     string `form:"name"     validate:"omitempty,min=2,max=100"`
	Email    string `form:"email"    validate:"omitempty,email,max=100"`
	Password string `form:"password" validate:"omitempty,min=8,max=255"`
}
