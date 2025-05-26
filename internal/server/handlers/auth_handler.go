package handlers

import (
	"log"
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
// @Param login body requests.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginSuccess
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
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
		log.Printf("error: %v", err)
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

// Register Doc
// @Summary Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param registerRequest body requests.RegisterRequest true "Register Request"
// @Success 200 {object} models.RegisterSuccess
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /api/v1/register [post]
func (ah *AuthHandler) Register(c echo.Context) error {
	registerRequest := new(r.RegisterRequest)
	if err := c.Bind(&registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Data is not valid")
	}

	if err := utils.ValidateRegisterRequest(*registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Data is not valid")
	}

	user, err := ah.authService.Register(*registerRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return responses.SuccessResponse(c, http.StatusOK, "Register Success", user.Email)

}

func (ah *AuthHandler) UpdatePassword(c echo.Context) error {
	return nil
}
