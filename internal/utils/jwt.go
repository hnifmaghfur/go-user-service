package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/models"
)

type JwtAccessClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user models.User, cfg config.AuthConfig) (string, error) {
	expiresIn, err := time.ParseDuration(cfg.AccessTokenExpiresIn)
	if err != nil {
		return "", err // or handle error appropriately
	}
	claims := JwtAccessClaims{
		Name: user.Name,
		ID:   user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.AccessTokenSecretKey))
}

func GenerateRefreshToken(user models.User, cfg config.AuthConfig) (string, error) {
	expiresIn, err := time.ParseDuration(cfg.RefreshTokenExpiresIn)
	if err != nil {
		return "", err // or handle error appropriately
	}
	claims := JwtRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.RefreshTokenSecretKey))
}
