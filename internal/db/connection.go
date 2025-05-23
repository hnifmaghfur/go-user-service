package db

import (
	"fmt"

	"github.com/hnifmaghfur/go-user-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return db, nil
}

func dns(cfg config.DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
}
