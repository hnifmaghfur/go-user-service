package services

import (
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/models"
	"github.com/hnifmaghfur/go-user-service/internal/repositories"
	req "github.com/hnifmaghfur/go-user-service/internal/requests"
	res "github.com/hnifmaghfur/go-user-service/internal/responses"
	"github.com/hnifmaghfur/go-user-service/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository repositories.AuthRepository
	userRepository repositories.UserRepository
	cfg            config.AuthConfig
	mc             *memcache.Client
}

func NewAuthService(authRepository repositories.AuthRepository, userRepository repositories.UserRepository, cfg config.AuthConfig, mc *memcache.Client) *AuthService {
	return &AuthService{authRepository: authRepository, userRepository: userRepository, cfg: cfg, mc: mc}
}

func (s *AuthService) Login(req req.LoginRequest) (res.TokenResponse, error) {
	// check user with email
	user, err := s.authRepository.Login(req.BasicAuth.Email)
	if err != nil {
		return res.TokenResponse{}, err
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.BasicAuth.Password)); err != nil {
		return res.TokenResponse{}, err
	}

	// config Access Token
	accessToken, err := utils.GenerateAccessToken(user, s.cfg)
	if err != nil {
		return res.TokenResponse{}, err
	}

	// config Refresh Token
	refreshToken, err := utils.GenerateRefreshToken(user, s.cfg, s.mc)
	if err != nil {
		return res.TokenResponse{}, err
	}

	return res.TokenResponse{
		LoginResponse: res.LoginResponse{
			AccessToken: accessToken,
			ExpiresIn:   s.cfg.AccessTokenExpiresIn,
		},
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Register(req req.RegisterRequest) (models.User, error) {
	// check email exist or not
	_, err := s.authRepository.Login(req.Email)
	if err == nil {
		return models.User{}, err
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.cfg.BcryptCost)
	log.Printf("hashed password: %q", hashedPassword)
	if err != nil {
		return models.User{}, err
	}

	// create user
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
		Phone:    req.Phone,
	}
	if err := s.userRepository.Post(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *AuthService) UpdateAccessToken(req req.UpdateAccessTokenRequest) (res.TokenResponse, error) {

	// check user exist
	user, err := s.authRepository.GetUserById(req.TokenId)
	if err != nil {
		return res.TokenResponse{}, err
	}

	// check refresh token
	item, err := s.mc.Get(fmt.Sprintf("refresh_token:%d", req.TokenId))
	if err != nil {
		return res.TokenResponse{}, err
	}

	// compare refresh token
	if err := bcrypt.CompareHashAndPassword(item.Value, []byte(req.RefreshToken)); err != nil {
		return res.TokenResponse{}, err
	}

	// generate new access token
	accessToken, err := utils.GenerateAccessToken(user, s.cfg)
	if err != nil {
		return res.TokenResponse{}, err
	}

	// generate new refresh token
	refreshToken, err := utils.GenerateRefreshToken(user, s.cfg, s.mc)
	if err != nil {
		return res.TokenResponse{}, err
	}

	return res.TokenResponse{
		LoginResponse: res.LoginResponse{
			AccessToken: accessToken,
			ExpiresIn:   s.cfg.AccessTokenExpiresIn,
		},
		RefreshToken: refreshToken,
	}, nil

}
