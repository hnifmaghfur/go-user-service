package handlers

import (
	"log"
	"net/http"

	"github.com/hnifmaghfur/go-user-service/internal/config"
	r "github.com/hnifmaghfur/go-user-service/internal/requests"
	"github.com/hnifmaghfur/go-user-service/internal/responses"
	"github.com/hnifmaghfur/go-user-service/internal/services"
	"github.com/hnifmaghfur/go-user-service/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	authService *services.AuthService
	cfg         config.AuthConfig
}

func NewAuthHandler(authService *services.AuthService, cfg config.AuthConfig) AuthHandler {
	return AuthHandler{authService: authService, cfg: cfg}
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

	err = ah.SetRefreshTokenCookie(c, token.RefreshToken)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

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

// Update Access Token Doc
// @Summary Update Access Token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} models.UpdateAccessTokenSuccess
// @Failure 400 {object} models.ErrorBadRequestResponse
// @Failure 500 {object} models.ErrorInternalServerErrorResponse
// @Router /api/v1/auth/update-token [post]
func (ah *AuthHandler) UpdateAccessToken(c echo.Context) error {
	// private func for remove cookies
	removeCookie := func() error {
		c.SetCookie(&http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		})
		return nil
	}

	if c.Get("token_id") == nil && c.Get("token") == nil {
		removeCookie()
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid Token")
	}

	reqUpdateAccessToken := new(r.UpdateAccessTokenRequest)
	reqUpdateAccessToken.TokenId = c.Get("token_id").(uint)
	reqUpdateAccessToken.RefreshToken = c.Get("token").(string)

	newToken, err := ah.authService.UpdateAccessToken(*reqUpdateAccessToken)
	if err != nil {
		removeCookie()
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	err = ah.SetRefreshTokenCookie(c, newToken.RefreshToken)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return responses.SuccessResponse(c, http.StatusCreated, "Update Access Token Success", newToken.LoginResponse)
}

func (ah *AuthHandler) GoogleLogin(c echo.Context) error {
	url := utils.NewGoogleConfig(ah.cfg).AuthCodeURL(ah.cfg.Google.State, oauth2.AccessTypeOffline)
	return responses.SuccessResponse(c, http.StatusOK, "Google Login Success", url)
}

func (ah *AuthHandler) GoogleCallback(c echo.Context) error {
	log.Println(c.QueryParam(ah.cfg.Google.State))
	return nil
}

func (ah *AuthHandler) SetRefreshTokenCookie(c echo.Context, token string) error {
	// insert refersh token on cookies
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
	return nil
}
