package main

import (
	"github.com/cubdesign/haikara/config"
	"github.com/cubdesign/haikara/controllers"
	"github.com/cubdesign/haikara/database"
	"github.com/cubdesign/haikara/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	database.Init()
	defer database.Close()
	controllers.StartWebserver()
}
