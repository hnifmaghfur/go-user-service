package routes

import (
	"github.com/hnifmaghfur/go-user-service/internal/repositories"
	s "github.com/hnifmaghfur/go-user-service/internal/server"
	handlers "github.com/hnifmaghfur/go-user-service/internal/server/handlers"
	"github.com/hnifmaghfur/go-user-service/internal/services"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRoutes(server *s.Server) error {
	authRepository := repositories.NewAuthRepository(server.DB)
	authService := services.NewAuthService(authRepository, server.Cfg.Auth)
	authHandler := handlers.NewAuthHandler(authService)

	// Prefix API
	api := server.Echo.Group("/api/v1")

	// swagger API
	api.GET("/swagger/*", echoSwagger.WrapHandler)

	// API Login
	api.POST("/login", authHandler.Login)

	// API Register
	api.POST("/register", authHandler.Register)

	return nil
}
