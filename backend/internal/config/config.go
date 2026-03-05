package config

import (
	"os"
	"time"
)

type Config struct {
	// Database
	DatabaseURL string

	// Server
	Port        string
	Host        string
	Environment string

	// Security
	JWTSecret      string
	JWTExpiration  time.Duration
	AgentSecretKey string
	AgentTimeout   time.Duration

	// Logging
	LogLevel string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL:    getEnv("DATABASE_URL", "postgresql://vps_user:vps_password@localhost:5432/vps_security"),
		Port:           getEnv("PORT", "8080"),
		Host:           getEnv("HOST", "0.0.0.0"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration:  parseDuration(getEnv("JWT_EXPIRATION", "24h")),
		AgentSecretKey: getEnv("AGENT_SECRET_KEY", "agent-secret-key"),
		AgentTimeout:   parseDuration(getEnv("AGENT_TIMEOUT", "30s")),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return time.Duration(0)
	}
	return duration
}
