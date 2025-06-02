package utils

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/models"
	"golang.org/x/crypto/bcrypt"
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
		return "", err
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

func GenerateRefreshToken(user models.User, cfg config.AuthConfig, mc *memcache.Client) (string, error) {
	expiresIn, err := time.ParseDuration(cfg.RefreshTokenExpiresIn)
	if err != nil {
		return "", err
	}
	claims := JwtRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSign, err := token.SignedString([]byte(cfg.RefreshTokenSecretKey))
	if err != nil {
		return "", err
	}

	encToken := tokenSign
	if len(encToken) > 72 {
		encToken = encToken[:72]
	}

	// hash token before store
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(encToken), cfg.BcryptCost)
	if err != nil {
		return "", err
	}

	// store refresh token on memcache
	mc.Set(&memcache.Item{
		Key:        fmt.Sprintf("refresh_token:%d", user.ID),
		Value:      hashedToken,
		Expiration: int32(expiresIn.Seconds()),
	})

	return tokenSign, nil
}

func VerifyRefreshToken(tokenString string, cfg config.AuthConfig) (*JwtRefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtRefreshClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.RefreshTokenSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token.Claims.(*JwtRefreshClaims), nil
}
