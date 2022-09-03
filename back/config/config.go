package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type ConfigList struct {
	LogFile string
	CORS    []string
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
		LogFile: os.Getenv("LOG_FILE"),
		CORS:    cors,
	}
}
