package model

import (
	"time"

	"agussyahrilmubarok.github.io/backend/internal/domain"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=100"`
	Email    string `json:"email" binding:"omitempty,email,max=100"`
	Password string `json:"password" binding:"omitempty,min=8,max=255"`
}
