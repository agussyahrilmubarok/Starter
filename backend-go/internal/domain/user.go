package domain

import (
	"errors"
	"time"
)

// User represents the user entity within the system.
type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserEmailExists      = errors.New("user email exists")
	ErrUserEmailNotFound    = errors.New("user email not found")
	ErrUserPasswordNotMatch = errors.New("user password not match")
)
