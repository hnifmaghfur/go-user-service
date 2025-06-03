package utils

import (
	"strings"

	"github.com/hnifmaghfur/go-user-service/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleConfig(cfg config.AuthConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.Google.ClientID,
		ClientSecret: cfg.Google.ClientSecret,
		RedirectURL:  cfg.Google.RedirectUrl,
		Scopes:       strings.Split(cfg.Google.ScopeUrl+cfg.Google.Scopes, ","),
		Endpoint:     google.Endpoint,
	}
}
