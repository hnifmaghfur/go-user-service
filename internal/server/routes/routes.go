package routes

import (
	_ "github.com/hnifmaghfur/go-user-service/docs"
	"github.com/hnifmaghfur/go-user-service/internal/middlewares"
	"github.com/hnifmaghfur/go-user-service/internal/repositories"
	s "github.com/hnifmaghfur/go-user-service/internal/server"
	handlers "github.com/hnifmaghfur/go-user-service/internal/server/handlers"
	"github.com/hnifmaghfur/go-user-service/internal/services"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRoutes(server *s.Server) error {
	authRepository := repositories.NewAuthRepository(server.DB)
	userRepository := repositories.NewUserRepository(server.DB)
	authService := services.NewAuthService(authRepository, userRepository, server.Cfg.Auth, server.Mc)
	authHandler := handlers.NewAuthHandler(authService)

	// Prefix API
	api := server.Echo.Group("/api/v1")

	// swagger API
	api.GET("/swagger/*", echoSwagger.WrapHandler)

	// apiCookie API
	apiCookie := api.Group("")
	apiCookie.Use(middlewares.CookieMiddleware(server.Cfg.Auth))

	// API Auth
	auth := api.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.POST("/register", authHandler.Register)

	// API Auth with Cookie
	authCookie := apiCookie.Group("/auth")
	authCookie.POST("/update-token", authHandler.UpdateAccessToken)

	return nil
}
