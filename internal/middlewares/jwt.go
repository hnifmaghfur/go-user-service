package middlewares

import (
	"net/http"
	"strings"

	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/responses"
	"github.com/hnifmaghfur/go-user-service/internal/utils"
	"github.com/labstack/echo/v4"
)

func CookieMiddleware(cfg config.AuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Cookie")
			if authHeader == "" {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Missing Cookie Header")
			}

			tokenString := authHeader[strings.Index(authHeader, "=")+1:]

			claim, err := utils.VerifyRefreshToken(tokenString, cfg)
			if err != nil {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid Token")
			}

			tokenId := claim.ID
			c.Set("token_id", tokenId)
			c.Set("token", tokenString)

			return next(c)
		}
	}
}
