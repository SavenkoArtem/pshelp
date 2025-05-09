package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("ERROR: .env is not loaded")
	}

	return &Config{
		Db: DbConfig{
			Dsn: getString("DSN", ""),
		},
	}

}

func getString(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	val, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return val
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	val, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return val
}
