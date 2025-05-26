package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/caarlos0/env/v11"
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/hnifmaghfur/go-user-service/internal/db"
	"github.com/hnifmaghfur/go-user-service/internal/server"
	"github.com/hnifmaghfur/go-user-service/internal/server/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title User Service
// @version 1.0
// @description This is user service application with Oauth2 and JWT security.
// @host localhost:${PORT}

// @contact.name Hanif Maghfur
// @contact.url https://maghfur.dev
// @contact.email hanif@maghfur.dev

// @BasePath /v2
func main() {
	// running service
	if err := run(); err != nil {
		slog.Error("Error running service:", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	// Parse .env file
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("failed to parse .env file: %w", err)
	}

	// Initialize database
	db, err := db.NewGormDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// Initialize memcache
	mc := memcache.New(cfg.Memcache.Host + ":" + cfg.Memcache.Port)

	// Set database connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Second)

	// Defer the close to clean up the connection pool when the application exits
	defer sqlDB.Close()

	// Initialize and route server
	app := server.NewServer(echo.New(), db, mc, cfg)
	routes.NewRoutes(app)

	// Start server
	if err := app.Start(fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
