package handlers

import (
	"net/http"

	r "github.com/hnifmaghfur/go-user-service/internal/requests"
	"github.com/hnifmaghfur/go-user-service/internal/responses"
	"github.com/hnifmaghfur/go-user-service/internal/services"
	"github.com/hnifmaghfur/go-user-service/internal/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) AuthHandler {
	return AuthHandler{authService: authService}
}

// Login Doc
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body requests.LoginRequest true "Login Request"
// @Success 200 {object} responses.TokenResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/login [post]
func (ah *AuthHandler) Login(c echo.Context) error {
	loginRequest := new(r.LoginRequest)
	if err := c.Bind(&loginRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Data is not valid")
	}

	if err := utils.ValidateLoginRequest(*loginRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Data is not valid")
	}

	token, err := ah.authService.Login(*loginRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	// insert refersh token on cookies
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	return responses.SuccessResponse(c, http.StatusOK, "Login Success", token.LoginResponse)
}

func (ah *AuthHandler) Register(c echo.Context) error {
	return nil
}

func (ah *AuthHandler) UpdatePassword(c echo.Context) error {
	return nil
}
