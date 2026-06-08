package usecase_test

import (
	"context"
	"testing"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/domain/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newUserUC(t *testing.T) (usecase.UserUseCase, *mocks.UserRepository) {
	repo := mocks.NewUserRepository(t)
	uc := usecase.NewUserUseCase(repo)
	return uc, repo
}

func stubUser() *domain.User {
	return &domain.User{
		ID:        uuid.New(),
		Name:      "Alice",
		Email:     "alice@mail.com",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestUser_GetAll(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocks.UserRepository)
		wantLen   int
		wantErr   bool
	}{
		{
			name: "success",
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindAll", mock.Anything).
					Return([]domain.User{*stubUser(), *stubUser()}, nil)
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "empty",
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindAll", mock.Anything).
					Return([]domain.User{}, nil)
			},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "repo error",
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindAll", mock.Anything).
					Return(nil, assert.AnError)
			},
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newUserUC(t)
			tt.mockSetup(repo)

			result, err := uc.GetAll(context.Background())

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, tt.wantLen)
			}
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	user := stubUser()

	tests := []struct {
		name      string
		id        uuid.UUID
		mockSetup func(*mocks.UserRepository)
		wantErr   error
	}{
		{
			name: "success",
			id:   user.ID,
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			id:   uuid.New(),
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: usecase.ErrUserNotFound,
		},
		{
			name: "repo error",
			id:   uuid.New(),
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByID", mock.Anything, mock.Anything).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newUserUC(t)
			tt.mockSetup(repo)

			result, err := uc.GetByID(context.Background(), tt.id)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, user.ID, result.ID)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	tests := []struct {
		name        string
		req         dto.CreateUserRequest
		mockSetup   func(*mocks.UserRepository)
		wantErr     error
		extraAssert func(*testing.T, *dto.UserResponse)
	}{
		{
			name: "success",
			req: dto.CreateUserRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.UserResponse) {
				assert.Equal(t, "Alice", res.Name)
				assert.Equal(t, "alice@mail.com", res.Email)
			},
		},
		{
			name: "email lowercased",
			req: dto.CreateUserRequest{
				Name:     "Alice",
				Email:    "Alice@Mail.COM",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.UserResponse) {
				assert.Equal(t, "alice@mail.com", res.Email)
			},
		},
		{
			name: "password is hashed",
			req: dto.CreateUserRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, mock.Anything).Return(false, nil)
				repo.On("Create", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					return u.Password != "secret123" && u.Password != ""
				})).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "email already exist",
			req: dto.CreateUserRequest{
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, mock.Anything).Return(true, nil)
			},
			wantErr: usecase.ErrEmailAlreadyInUse,
		},
		{
			name: "exists by email error",
			req: dto.CreateUserRequest{
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, mock.Anything).Return(false, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "repo create error",
			req: dto.CreateUserRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, mock.Anything).Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newUserUC(t)
			tt.mockSetup(repo)

			result, err := uc.Create(context.Background(), tt.req)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if tt.extraAssert != nil {
					tt.extraAssert(t, result)
				}
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(*mocks.UserRepository, *domain.User)
		req         dto.UpdateUserRequest
		wantErr     error
		extraAssert func(*testing.T, *dto.UserResponse)
	}{
		{
			name: "success name only",
			req:  dto.UpdateUserRequest{Name: "Alice Updated", Email: "alice@mail.com"},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.UserResponse) {
				assert.Equal(t, "Alice Updated", res.Name)
			},
		},
		{
			name: "success email changed",
			req:  dto.UpdateUserRequest{Name: "Alice", Email: "new@mail.com"},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("ExistsByEmail", mock.Anything, "new@mail.com").Return(false, nil)
				repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.UserResponse) {
				assert.Equal(t, "new@mail.com", res.Email)
			},
		},
		{
			name: "success password changed",
			req:  dto.UpdateUserRequest{Name: "Alice", Email: "alice@mail.com", Password: "newpassword123"},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				oldPassword := user.Password
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("Update", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					return u.Password != oldPassword && u.Password != "newpassword123"
				})).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			req:  dto.UpdateUserRequest{},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(nil, nil)
			},
			wantErr: usecase.ErrUserNotFound,
		},
		{
			name: "email conflict",
			req:  dto.UpdateUserRequest{Name: "Alice", Email: "taken@mail.com"},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("ExistsByEmail", mock.Anything, "taken@mail.com").Return(true, nil)
			},
			wantErr: usecase.ErrEmailAlreadyInUse,
		},
		{
			name: "find by id error",
			req:  dto.UpdateUserRequest{},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "repo update error",
			req:  dto.UpdateUserRequest{Name: "Alice", Email: "alice@mail.com"},
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("Update", mock.Anything, mock.AnythingOfType("*domain.User")).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newUserUC(t)
			user := stubUser()
			tt.mockSetup(repo, user)

			result, err := uc.Update(context.Background(), user.ID, tt.req)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if tt.extraAssert != nil {
					tt.extraAssert(t, result)
				}
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocks.UserRepository, *domain.User)
		wantErr   error
	}{
		{
			name: "success",
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("Delete", mock.Anything, user.ID).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "not found",
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(nil, nil)
			},
			wantErr: usecase.ErrUserNotFound,
		},
		{
			name: "find by id error",
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "repo delete error",
			mockSetup: func(repo *mocks.UserRepository, user *domain.User) {
				repo.On("FindByID", mock.Anything, user.ID).Return(user, nil)
				repo.On("Delete", mock.Anything, user.ID).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newUserUC(t)
			user := stubUser()
			tt.mockSetup(repo, user)

			err := uc.Delete(context.Background(), user.ID)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
