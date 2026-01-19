package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort string
	DbPath  string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("ENV file not loaded")
	}

	e.AppPort = loadString("APP_PORT", ":8080")
	e.DbPath = loadString("DB_PATH", "postgres://postgres:adminPassword@localhost:5433/tasks?sslmode=disable")
}

var Config envConfig

func init() {
	Config.LoadConfig()
}

func loadString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("Key %s is not loaded", key)
		return fallback
	}
	return val
}
