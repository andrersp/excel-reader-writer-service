package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_ENV, APP_PORT string
)

func LoadConfig() error {

	err := godotenv.Load()
	APP_ENV = getEnv("APP_ENV", "TEST")

	if err != nil && APP_ENV == "LOCAL" {
		return err

	}
	APP_PORT = getEnv("APP_PORT", "8080")

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}
	return value
}
