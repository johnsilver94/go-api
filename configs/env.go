package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	DbUrl                  string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		Port:                   getEnv("PORT", "8080"),
		DbUrl:                  getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		JWTSecret:              getEnv("JWT_SECRET", "secret"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_ACCESS_TOKEN_EXPIRATION_TIME", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
