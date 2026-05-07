package dto

import "backend/internal/domain"

type UserDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (ud *UserDTO) FromDomain(user *domain.User) {
	if user == nil {
		ud = nil
	}

	ud.ID = user.ID
	ud.Name = user.Name
	ud.Email = user.Email
	ud.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	ud.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password,omitempty"`
}
