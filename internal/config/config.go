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
}

type Memcache struct {
	Host string `env:"MEMCACHE_HOST"`
	Port string `env:"MEMCACHE_PORT"`
}
