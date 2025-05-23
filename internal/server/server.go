package server

import (
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
	Cfg  config.Config
}

func NewServer(echo *echo.Echo, db *gorm.DB, cfg config.Config) *Server {
	return &Server{
		Echo: echo,
		DB:   db,
		Cfg:  cfg,
	}
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}
