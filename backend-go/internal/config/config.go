package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
	JWTSecret  string
	JWTExpiry  time.Duration
}

func LoadConfig() (*Config, error) {
	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "aicg"),
		DBPort:     getEnv("DB_PORT", "5432"),
		ServerPort: getEnv("PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
	}

	// Parse JWT expiration
	expiryStr := getEnv("JWT_EXPIRATION", "24h")
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		return nil, err
	}
	config.JWTExpiry = expiry

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
