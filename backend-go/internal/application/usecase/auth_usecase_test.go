package usecase_test

import (
	"context"
	"testing"

	"agussyahrilmubarok.github.io/backend/internal/application/dto"
	"agussyahrilmubarok.github.io/backend/internal/application/usecase"
	"agussyahrilmubarok.github.io/backend/internal/domain"
	"agussyahrilmubarok.github.io/backend/internal/domain/mocks"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"agussyahrilmubarok.github.io/backend/pkg/crypto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newAuthUC(t *testing.T) (usecase.AuthUseCase, *mocks.UserRepository) {
	repo := mocks.NewUserRepository(t)
	jwt := security.NewJwtManager(&config.JWT{
		Secret:     "test-secret-key",
		ExpiryHour: 24,
	})
	uc := usecase.NewAuthUseCase(repo, jwt)
	return uc, repo
}

func TestAuth_SignUp(t *testing.T) {
	tests := []struct {
		name        string
		req         dto.SignUpRequest
		mockSetup   func(*mocks.UserRepository)
		wantErr     error
		extraAssert func(*testing.T, *dto.AuthResponse)
	}{
		{
			name: "success",
			req: dto.SignUpRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.AuthResponse) {
				assert.NotEmpty(t, res.Token)
				assert.Equal(t, "Alice", res.User.Name)
				assert.Equal(t, "alice@mail.com", res.User.Email)
			},
		},
		{
			name: "email lowercased",
			req: dto.SignUpRequest{
				Name:     "Alice",
				Email:    "Alice@Mail.COM",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "Alice@Mail.COM").Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.AuthResponse) {
				assert.Equal(t, "alice@mail.com", res.User.Email)
			},
		},
		{
			name: "email already exist",
			req: dto.SignUpRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(true, nil)
			},
			wantErr: usecase.ErrEmailAlreadyExist,
		},
		{
			name: "exists by email error",
			req: dto.SignUpRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(false, assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "repo create error",
			req: dto.SignUpRequest{
				Name:     "Alice",
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("ExistsByEmail", mock.Anything, "alice@mail.com").Return(false, nil)
				repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newAuthUC(t)
			tt.mockSetup(repo)

			result, err := uc.SignUp(context.Background(), tt.req)

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

func TestAuth_SignIn(t *testing.T) {
	hashed, _ := crypto.HashPassword("secret123")
	user := &domain.User{
		ID:       uuid.New(),
		Name:     "Alice",
		Email:    "alice@mail.com",
		Password: hashed,
	}

	tests := []struct {
		name        string
		req         dto.SignInRequest
		mockSetup   func(*mocks.UserRepository)
		wantErr     error
		extraAssert func(*testing.T, *dto.AuthResponse)
	}{
		{
			name: "success",
			req: dto.SignInRequest{
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByEmail", mock.Anything, "alice@mail.com").Return(user, nil)
			},
			wantErr: nil,
			extraAssert: func(t *testing.T, res *dto.AuthResponse) {
				assert.NotEmpty(t, res.Token)
				assert.Equal(t, user.ID, res.User.ID)
				assert.Equal(t, user.Email, res.User.Email)
			},
		},
		{
			name: "email not found",
			req: dto.SignInRequest{
				Email:    "notfound@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByEmail", mock.Anything, "notfound@mail.com").Return(nil, nil)
			},
			wantErr: usecase.ErrInvalidEmail,
		},
		{
			name: "wrong password",
			req: dto.SignInRequest{
				Email:    "alice@mail.com",
				Password: "wrongpassword",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByEmail", mock.Anything, "alice@mail.com").Return(user, nil)
			},
			wantErr: usecase.ErrInvalidPassword,
		},
		{
			name: "repo error",
			req: dto.SignInRequest{
				Email:    "alice@mail.com",
				Password: "secret123",
			},
			mockSetup: func(repo *mocks.UserRepository) {
				repo.On("FindByEmail", mock.Anything, "alice@mail.com").Return(nil, assert.AnError)
			},
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc, repo := newAuthUC(t)
			tt.mockSetup(repo)

			result, err := uc.SignIn(context.Background(), tt.req)

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
