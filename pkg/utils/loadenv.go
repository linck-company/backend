package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var loaded bool

// LoadEnv loads env var only once when the program execution starts
func LoadEnv() error {
	if !loaded {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, relying on system environment variables.")
		}
		loaded = true
	}
	return nil
}

func GetStringEnv(envKey, fallback string) string {
	if val, exists := os.LookupEnv(envKey); exists {
		return val
	}
	return fallback
}

func GetIntEnv(envKey string, fallback int) int {
	if val, exists := os.LookupEnv(envKey); exists {
		result, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return result
	}
	return fallback
}

func GetBoolEnv(envKey string, fallback bool) bool {
	if val, exists := os.LookupEnv(envKey); exists {
		result, err := strconv.ParseBool(val)
		if err != nil {
			return fallback
		}
		return result
	}
	return fallback
}
