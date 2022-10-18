package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
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

	PageSize int

	ChromeDevToolsHostAndPort string

	UserAgent string

	AssetsUrl string
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

	pageSize, err := strconv.Atoi(os.Getenv("PAGE_SIZE"))
	if err != nil {
		log.Fatalf("action: Load env page size format,  err: %s", err.Error())
	}

	Config = ConfigList{
		LogFile:                   os.Getenv("LOG_FILE"),
		DbHost:                    os.Getenv("DB_HOST"),
		DbName:                    os.Getenv("DB_NAME"),
		DbUser:                    os.Getenv("DB_USER"),
		DbPassword:                os.Getenv("DB_PASSWORD"),
		Port:                      os.Getenv("PORT"),
		CORS:                      cors,
		PageSize:                  pageSize,
		ChromeDevToolsHostAndPort: os.Getenv("CHROME_DEVTOOLS_HOST_AND_PORT"),
		UserAgent:                 os.Getenv("USER_AGENT"),
		AssetsUrl:                 os.Getenv("ASSETS_URL"),
	}
}
