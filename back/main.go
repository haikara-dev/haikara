package main

import (
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/controllers"
	"github.com/haikara-dev/haikara/database"
	"github.com/haikara-dev/haikara/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	database.Init()
	defer database.Close()
	controllers.StartWebserver()
}
