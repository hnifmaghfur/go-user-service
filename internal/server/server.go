package server

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/hnifmaghfur/go-user-service/internal/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
	Cfg  config.Config
	Mc   *memcache.Client
}

func NewServer(echo *echo.Echo, db *gorm.DB, mc *memcache.Client, cfg config.Config) *Server {
	return &Server{
		Echo: echo,
		DB:   db,
		Cfg:  cfg,
		Mc:   mc,
	}
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(addr)
}
