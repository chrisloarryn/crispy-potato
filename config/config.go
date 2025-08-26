package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
	Storage  StorageConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URI string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// StorageConfig holds storage configuration
type StorageConfig struct {
	BasePath string
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			URI: getEnv("MONGODB_URI", "mongodb://localhost:27017/twittor"),
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET", "MastersOfDevelopment_facebookGroup"),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Storage: StorageConfig{
			BasePath: getEnv("STORAGE_PATH", "."),
		},
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
