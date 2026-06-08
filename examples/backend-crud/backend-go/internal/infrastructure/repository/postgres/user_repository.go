package postgres

import (
	"context"
	"errors"

	"agussyahrilmubarok.github.io/backend/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// FindAll implements [domain.UserRepository].
func (r *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID implements [domain.UserRepository].
func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail implements [domain.UserRepository].
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Create implements [domain.UserRepository].
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Update implements [domain.UserRepository].
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete implements [domain.UserRepository].
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error
}

// ExistsByEmail implements [domain.UserRepository].
func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// Count implements [domain.UserRepository].
func (r *userRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
