package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type config struct {
	Env       string
	Addr      string
	LogLevel  string
	LogFormat string

	DatabaseUrl string
	DBMaxConns  int
	DBMinConns  int
	PingTimeout time.Duration

	BcryptCost int
	JWTSecret  string
}

func New(path string) (*config, error) {
	if err := loadEnv(path); err != nil {
		return nil, err
	}

	c := config{
		Env:       envString("ENV", "development"),
		Addr:      envString("HOST", "8080"),
		LogLevel:  envString("LOG_LEVEL", "info"),
		LogFormat: envString("LOG_FORMAT", "text"),

		DatabaseUrl: envString("DATABASE_URL", ""),
		DBMaxConns:  envInt("DB_MAX_CONN", 10),
		DBMinConns:  envInt("DB_MIN_CONN", 5),
		PingTimeout: envDur("PING_TIMEOUT", 3*time.Second),

		BcryptCost: envInt("BCRYPT_COST", 12),
		JWTSecret:  envString("JWT_SECRET", ""),
	}

	if err := c.Validate(); err != nil {
		return &config{}, err
	}

	return &c, nil
}

func oneOf(v string, allowed ...string) bool {
	for _, a := range allowed {
		if strings.EqualFold(v, a) {
			return true
		}
	}
	return false
}

func (c *config) Validate() error {

	if !oneOf(c.Env, "development", "staging", "production") {
		return fmt.Errorf("ENVIRONMENT must be one of: development, staging, production, got %s", c.Env)
	}
	if !oneOf(c.LogLevel, "debug", "info", "warn", "error") {
		return fmt.Errorf("LOG_LEVEL invalid: %q", c.LogLevel)
	}
	if !oneOf(c.LogFormat, "json", "text") {
		return fmt.Errorf("LOG_FORMAT invalid: %q", c.LogFormat)
	}
	if p, err := strconv.Atoi(c.Addr); err != nil || p < 1 || p > 65535 {
		return fmt.Errorf("HTTP_PORT invalid: %q", c.Addr)
	}
	if strings.TrimSpace(c.DatabaseUrl) == "" {
		return fmt.Errorf("DATABASE_URL is required")
	}
	if c.DBMaxConns < 1 {
		return fmt.Errorf("DB_MAX_CONNS must be >= 1, got %d", c.DBMaxConns)
	}
	if c.DBMinConns < 0 || c.DBMinConns > c.DBMaxConns {
		return fmt.Errorf("DB_MIN_CONNS must be 0..DB_MAX_CONNS (%d), got %d", c.DBMaxConns, c.DBMinConns)
	}

	return nil
}
