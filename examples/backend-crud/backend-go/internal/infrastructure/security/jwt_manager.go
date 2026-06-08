package security

import (
	"errors"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTManager interface {
	GenerateToken(id uuid.UUID) (string, error)
	ValidateToken(token string) (uuid.UUID, error)
}

type jwtManager struct {
	secret     string
	expiryHour int
}

func NewJwtManager(cfg *config.JWT) JWTManager {
	return &jwtManager{
		secret:     cfg.Secret,
		expiryHour: cfg.ExpiryHour,
	}
}

func (m *jwtManager) GenerateToken(id uuid.UUID) (string, error) {
	c := claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(m.expiryHour) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(m.secret))
}

func (m *jwtManager) ValidateToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return uuid.Nil, ErrExpiredToken
		}
		return uuid.Nil, ErrInvalidToken
	}

	return token.Claims.(*claims).UserID, nil
}
