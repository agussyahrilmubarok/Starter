package service

import (
	"context"
	"errors"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockery --name=IJWTService
type IJWTService interface {
	Generate(ctx context.Context, userID string) (string, error)
	Validate(ctx context.Context, tokenString string) (string, error)
}

type jwtService struct {
	config *config.JWTConfig
}

type jwtClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate implements [IJWTService].
func (s *jwtService) Generate(ctx context.Context, userID string) (string, error) {
	expTime, err := time.ParseDuration(s.config.ExpTime)
	if err != nil {
		return "", err
	}

	claims := jwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

// Validate implements [IJWTService].
func (s *jwtService) Validate(ctx context.Context, tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}

func NewJWTService(
	config *config.JWTConfig,
) IJWTService {
	return &jwtService{
		config: config,
	}
}
