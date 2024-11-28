package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DbUrl string
}

var Envs = initConfig()

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		Port:  getEnv("PORT", "8080"),
		DbUrl: getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
