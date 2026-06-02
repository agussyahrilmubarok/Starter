package crypto_test

import (
	"testing"

	"agussyahrilmubarok.github.io/backend/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "success",
			password: "secret123",
			wantErr:  false,
		},
		{
			name:     "minimum length",
			password: "12345678",
			wantErr:  false,
		},
		{
			name:     "special characters",
			password: "p@ssw0rd!#$",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := crypto.HashPassword(tt.password)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, hashed)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, hashed)
				assert.NotEqual(t, tt.password, hashed)
			}
		})
	}
}

func TestHashPassword_IsUnique(t *testing.T) {
	hashed1, err := crypto.HashPassword("secret123")
	assert.NoError(t, err)

	hashed2, err := crypto.HashPassword("secret123")
	assert.NoError(t, err)

	assert.NotEqual(t, hashed1, hashed2)
}

func TestCheckPassword(t *testing.T) {
	plain := "secret123"
	hashed, err := crypto.HashPassword(plain)
	assert.NoError(t, err)

	tests := []struct {
		name   string
		hashed string
		plain  string
		want   bool
	}{
		{
			name:   "correct password",
			hashed: hashed,
			plain:  plain,
			want:   true,
		},
		{
			name:   "wrong password",
			hashed: hashed,
			plain:  "wrongpassword",
			want:   false,
		},
		{
			name:   "empty plain",
			hashed: hashed,
			plain:  "",
			want:   false,
		},
		{
			name:   "empty hashed",
			hashed: "",
			plain:  plain,
			want:   false,
		},
		{
			name:   "both empty",
			hashed: "",
			plain:  "",
			want:   false,
		},
		{
			name:   "invalid hash format",
			hashed: "not-a-bcrypt-hash",
			plain:  plain,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := crypto.CheckPassword(tt.hashed, tt.plain)
			assert.Equal(t, tt.want, result)
		})
	}
}
