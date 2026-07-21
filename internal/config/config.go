package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		// Running without .env is acceptable.
		// Environment variables may already be set.
		log.Println("No .env file found; using environment variables")

	}
	cfg := &Config{
		AppPort:      getEnv("APP_PORT", "9000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", ""),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", ""),
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
