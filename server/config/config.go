package config

import (
	"os"
	"time"
)

type Config struct {
	JWTAccessSecret    string
	JWTRefreshSecret   string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

func LoadConfig() *Config {
	return &Config{
		JWTAccessSecret:    getEnv("JWT_ACCESS_SECRET", "your-access-secret-key-change-in-production"),
		JWTRefreshSecret:   getEnv("JWT_REFRESH_SECRET", "your-refresh-secret-key-change-in-production"),
		AccessTokenExpiry:  5 * time.Minute,    // 5 minutes
		RefreshTokenExpiry: 7 * 24 * time.Hour, // 7 days
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
