package services

import (
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/repositories"
	req "github.com/hnifmaghfur/go-user-service/internal/requests"
	res "github.com/hnifmaghfur/go-user-service/internal/responses"
	"github.com/hnifmaghfur/go-user-service/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository repositories.AuthRepository
	cfg            config.AuthConfig
}

func NewAuthService(authRepository repositories.AuthRepository, cfg config.AuthConfig) *AuthService {
	return &AuthService{authRepository: authRepository, cfg: cfg}
}

func (s *AuthService) Login(req req.LoginRequest) (res.TokenResponse, error) {
	// check user with email
	user, err := s.authRepository.Login(req.BasicAuth)
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
	refreshToken, err := utils.GenerateRefreshToken(user, s.cfg)
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

func (s *AuthService) Register() error {
	return nil
}

func (s *AuthService) UpdatePassword() error {
	return nil
}
