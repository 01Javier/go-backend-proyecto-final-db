package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBService  string
	ServerPort string
	JWTSecret  string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found!")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "system"),
		DBPassword: getEnv("DB_PASSWORD", "oracle"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "1521"),
		DBService:  getEnv("DB_SERVICE", "XEPDB1"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
