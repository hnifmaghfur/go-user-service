package config

type Config struct {
	HTTP     HTTPConfig
	DB       DBConfig
	Auth     AuthConfig
	Memcache Memcache
}

type HTTPConfig struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
}

type DBConfig struct {
	User            string `env:"DB_USER"`
	Password        string `env:"DB_PASSWORD"`
	Driver          string `env:"DB_DRIVER"`
	Name            string `env:"DB_NAME"`
	Host            string `env:"DB_HOST"`
	Port            string `env:"DB_PORT"`
	MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS"`
	ConnMaxLifetime int    `env:"DB_CONN_MAX_LIFETIME"`
}

type AuthConfig struct {
	AccessTokenSecretKey  string `env:"ACCESS_TOKEN_SECRET_KEY"`
	AccessTokenExpiresIn  string `env:"ACCESS_TOKEN_EXPIRES_IN"`
	RefreshTokenSecretKey string `env:"REFRESH_TOKEN_SECRET_KEY"`
	RefreshTokenExpiresIn string `env:"REFRESH_TOKEN_EXPIRES_IN"`
	BcryptCost            int    `env:"BCRYPT_COST"`
	Google                GoogleConfig
}

type GoogleConfig struct {
	ClientID     string `env:"GOOGLE_CLIENT_ID"`
	ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	RedirectUrl  string `env:"GOOGLE_REDIRECT_URL"`
	Url          string `env:"GOOGLE_URL"`
	ScopeUrl     string `env:"GOOGLE_SCOPE_URL"`
	Scopes       string `env:"GOOGLE_SCOPES"`
	State        string `env:"GOOGLE_STATE"`
}

type Memcache struct {
	Host string `env:"MEMCACHE_HOST"`
	Port string `env:"MEMCACHE_PORT"`
}
