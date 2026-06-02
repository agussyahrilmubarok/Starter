package dto

import (
	"time"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string `json:"name"     binding:"required,min=2,max=100"`
	Email    string `json:"email"    binding:"required,email,max=150"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"     binding:"omitempty,min=2,max=100"`
	Email    string `json:"email"    binding:"omitempty,email,max=150"`
	Password string `json:"password" binding:"omitempty,min=8,max=72"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ur *UserResponse) FromDomain(user *domain.User) {
	ur.ID = user.ID
	ur.Name = user.Name
	ur.Email = user.Email
	ur.CreatedAt = user.CreatedAt
	ur.UpdatedAt = user.UpdatedAt
}

func NewUserResponse(user *domain.User) UserResponse {
	var ur UserResponse
	ur.FromDomain(user)
	return ur
}

func NewUserListResponse(users []domain.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i := range users {
		result[i].FromDomain(&users[i])
	}
	return result
}
