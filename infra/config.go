package infra

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	AppName          *string
	ENV              *string
	Port             *string

	DBHost     *string
	DBPort     *string
	DBName     *string
	DBUser     *string
	DBPassword *string
}

func Load() (*ServerConfig, error) {
	_ = godotenv.Load()

	return &ServerConfig{
		AppName:          getString("APP_NAME"),
		ENV:              getString("ENV"),
		Port:             getString("PORT"),

		DBHost:     getString("DB_HOST"),
		DBPort:     getString("DB_PORT"),
		DBName:     getString("DB_NAME"),
		DBUser:     getString("DB_USER"),
		DBPassword: getString("DB_PASSWORD"),

	}, nil
}

func getString(key string) *string {
	value := os.Getenv(key)

	return &value
}

func getInt(key string) (*int, error) {
	value := os.Getenv(key)

	if value == "" {
		return nil, fmt.Errorf("%s is required", key)
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s=%q: %w",
			key,
			value,
			err,
		)
	}

	return &result, nil
}

func getDuration(key string) (*time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return nil, fmt.Errorf("%s is required", key)
	}

	result, err := time.ParseDuration(value)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid %s=%q: %w",
			key,
			value,
			err,
		)
	}

	return &result, nil
}

func getBytes(key string) []byte {
	value := os.Getenv(key)

	return []byte(value)
}
