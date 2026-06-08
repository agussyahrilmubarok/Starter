package security_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"agussyahrilmubarok.github.io/backend/internal/infrastructure/security"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newJwtManager() security.JWTManager {
	return security.NewJwtManager(&config.JWT{
		Secret:     "test-secret-key",
		ExpiryHour: 24,
	})
}

func TestGenerateToken(t *testing.T) {
	m := newJwtManager()
	userID := uuid.New()

	token, err := m.GenerateToken(userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	m := newJwtManager()
	userID := uuid.New()

	token, err := m.GenerateToken(userID)
	assert.NoError(t, err)

	gotID, err := m.ValidateToken(token)

	assert.NoError(t, err)
	assert.Equal(t, userID, gotID)
}

func TestValidateToken_Invalid(t *testing.T) {
	m := newJwtManager()

	tests := []struct {
		name    string
		token   string
		wantErr error
	}{
		{
			name:    "empty token",
			token:   "",
			wantErr: security.ErrInvalidToken,
		},
		{
			name:    "random string",
			token:   "not.a.token",
			wantErr: security.ErrInvalidToken,
		},
		{
			name: "wrong secret",
			token: func() string {
				other := security.NewJwtManager(&config.JWT{
					Secret:     "wrong-secret",
					ExpiryHour: 24,
				})
				token, _ := other.GenerateToken(uuid.New())
				return token
			}(),
			wantErr: security.ErrInvalidToken,
		},
		{
			name: "invalid signing method",
			token: func() string {
				key, _ := rsa.GenerateKey(rand.Reader, 2048)
				c := jwt.MapClaims{
					"user_id": uuid.New().String(),
					"exp":     time.Now().Add(24 * time.Hour).Unix(),
					"iat":     time.Now().Unix(),
				}
				token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
				signed, _ := token.SignedString(key)
				return signed
			}(),
			wantErr: security.ErrInvalidToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := m.ValidateToken(tt.token)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, uuid.Nil, gotID)
		})
	}
}

func TestValidateToken_Expired(t *testing.T) {
	m := security.NewJwtManager(&config.JWT{
		Secret:     "test-secret-key",
		ExpiryHour: 0,
	})

	token, err := m.GenerateToken(uuid.New())
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	gotID, err := m.ValidateToken(token)

	assert.ErrorIs(t, err, security.ErrExpiredToken)
	assert.Equal(t, uuid.Nil, gotID)
}
