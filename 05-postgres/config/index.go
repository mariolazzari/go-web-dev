package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Env file not loaded")
	}

	e.AppPort = loadString("APP_PORT", "8080")
}

var Config envConfig

func loadString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Missing var: %s", key)
		return fallback
	}

	return val
}

func init() {
	Config.LoadConfig()
}
