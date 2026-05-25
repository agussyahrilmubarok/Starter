package service

import (
	"context"
	"errors"
	"time"

	"agussyahrilmubarok.github.io/backend/internal/config"
	"agussyahrilmubarok.github.io/backend/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
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
	log := logger.FromCtx(ctx).With(zap.String("user_id", userID))

	expTime, err := time.ParseDuration(s.config.ExpTime)
	if err != nil {
		log.Error("failed to parse exp time", zap.Error(err))
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
	signed, err := token.SignedString([]byte(s.config.SecretKey))
	if err != nil {
		log.Error("failed to sign token", zap.Error(err))
		return "", err
	}

	log.Debug("token generated", zap.Time("expires_at", claims.ExpiresAt.Time))
	return signed, nil
}

// Validate implements [IJWTService].
func (s *jwtService) Validate(ctx context.Context, tokenString string) (string, error) {
	log := logger.FromCtx(ctx)

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		log.Warn("failed to parse token", zap.Error(err))
		return "", err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok || !token.Valid {
		log.Warn("invalid token claims")
		return "", errors.New("invalid token")
	}

	log.Debug("token valid", zap.String("user_id", claims.UserID))
	return claims.UserID, nil
}

func NewJWTService(config *config.JWTConfig) IJWTService {
	return &jwtService{
		config: config,
	}
}
