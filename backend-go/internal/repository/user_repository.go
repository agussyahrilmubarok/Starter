package repository

import (
	"backend/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockery --name=IUserRepository
type IUserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, userID uint) (*domain.User, error)
	FindByEmail(ctx context.Context, userEmail string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Save(ctx context.Context, user *domain.User) error
	DeleteByID(ctx context.Context, userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

// FindAll implements [IUserRepository].
func (r *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).
		Find(&users).
		Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID implements [IUserRepository].
func (r *userRepository) FindByID(ctx context.Context, userID uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).
		First(&user, userID).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail implements [IUserRepository].
func (r *userRepository) FindByEmail(ctx context.Context, userEmail string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).
		Where("email = ?", userEmail).
		First(&user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Create implements [IUserRepository].
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).
		Create(&user).
		Error; err != nil {
		return err
	}
	return nil
}

// Save implements [IUserRepository].
func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).
		Save(&user).
		Error; err != nil {
		return err
	}
	return nil
}

// DeleteByID implements [IUserRepository].
func (r *userRepository) DeleteByID(ctx context.Context, userID uint) error {
	if err := r.db.WithContext(ctx).
		Delete(&domain.User{}, userID).
		Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}
