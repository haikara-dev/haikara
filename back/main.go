package main

import (
	"fmt"
	"github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/controllers"
	"github.com/haikara-dev/haikara/database"
	"github.com/haikara-dev/haikara/libs"
	"github.com/haikara-dev/haikara/utils"
	"time"
)

func ogp_image() {
	fmt.Println("OGP Image...")
	for {
		time.Sleep(10 * time.Minute) // 時間	10分休む
		// ここでなにかを行う
		libs.CrawlOGPImage(100, database.Client)
	}
}
func index() {
	fmt.Println("Indexing...")
	for {
		time.Sleep(30 * time.Minute) // 時間	30分休む
		// ここでなにかを行う
		libs.IndexAllFeed(database.Client)
	}
}
func crawl() {
	fmt.Println("Crawling...")
	for {
		time.Sleep(3 * time.Hour) // 3時間	休む
		// ここでなにかを行う
		libs.CrawlAllSite(database.Client)
	}
}
func main() {
	utils.LoggingSettings(config.Config.LogFile)
	database.Init()
	defer database.Close()

	go crawl()
	go index()
	go ogp_image()

	controllers.StartWebserver()

}
