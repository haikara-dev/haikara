package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type ConfigList struct {
	LogFile string

	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string

	Port string
	CORS []string
}

var Config ConfigList

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("action: Load env file, err: %s", err.Error())
	}

	cors := strings.Split(os.Getenv("CORS"), ",")
	for i := range cors {
		cors[i] = strings.TrimSpace(cors[i])
	}

	Config = ConfigList{
		LogFile:    os.Getenv("LOG_FILE"),
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		Port:       os.Getenv("PORT"),
		CORS:       cors,
	}
}
