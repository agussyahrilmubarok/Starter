package payload

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

func ToUserListResponse(users []domain.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i, user := range users {
		result[i] = ToUserResponse(&user)
	}
	return result
}

type CreateUserRequest struct {
	Name     string `form:"name"     binding:"required,min=2,max=100"`
	Email    string `form:"email"    binding:"required,email,max=100"`
	Password string `form:"password" binding:"required,min=8,max=255"`
}

type UpdateUserRequest struct {
	Name     string `form:"name"     binding:"omitempty,min=2,max=100"`
	Email    string `form:"email"    binding:"omitempty,email,max=100"`
	Password string `form:"password" binding:"omitempty,min=8,max=255"`
}
