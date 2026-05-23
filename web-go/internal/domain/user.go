package domain

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserEmailExists      = errors.New("user email exists")
	ErrUserEmailNotFound    = errors.New("user email not found")
	ErrUserPasswordNotMatch = errors.New("user password not match")
)
