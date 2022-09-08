package controllers

import (
	"github.com/cubdesign/dailyfj/api"
	todoConfig "github.com/cubdesign/dailyfj/config"
	"github.com/cubdesign/dailyfj/database"
	"github.com/cubdesign/dailyfj/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func healthCheck(c *gin.Context) {
	log.Printf("ClientIP: %s\n", c.ClientIP())
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func StartWebserver() {

	apiUserHandler := api.UserHandler{
		Client: database.Client,
	}

	apiSiteHandler := api.SiteHandler{
		Client: database.Client,
	}

	apiArticleHandler := api.ArticleHandler{
		Client: database.Client,
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowOrigins = todoConfig.Config.CORS

	r.Use(cors.New(config))

	r.GET("/health-check", healthCheck)

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/users/create", apiUserHandler.CreateUserByUUIDAndEmail)
		authorized.GET("/users/current", apiUserHandler.GetCurrentUser)

		authorized.GET("/users", apiUserHandler.GetAllUsers)
		authorized.POST("/users", apiUserHandler.CreateUser)
		authorized.GET("/users/:id", apiUserHandler.GetUser)
		authorized.PUT("/users/:id", apiUserHandler.UpdateUser)
		authorized.DELETE("/users/:id", apiUserHandler.DeleteUser)

		authorized.GET("/sites", apiSiteHandler.GetAllSites)
		authorized.POST("/sites", apiSiteHandler.CreateSite)
		authorized.GET("/sites/:id", apiSiteHandler.GetSite)
		authorized.PUT("/sites/:id", apiSiteHandler.UpdateSite)
		authorized.DELETE("/sites/:id", apiSiteHandler.DeleteSite)
		authorized.PATCH("/sites/active/:id", apiSiteHandler.ActiveSite)
		authorized.PATCH("/sites/deActive/:id", apiSiteHandler.DeActiveSite)
		authorized.GET("/sites/run-crawling/:id", apiSiteHandler.RunCrawling)
		authorized.GET("/sites/get-rss-url/:id", apiSiteHandler.GetRssUrlBySiteId)
		authorized.GET("/sites/get-rss-url-by-url", apiSiteHandler.GetRssUrlByUrl)

		authorized.GET("/articles", apiArticleHandler.GetAllArticles)
		authorized.POST("/articles", apiArticleHandler.CreateArticle)
		authorized.GET("/articles/:id", apiArticleHandler.GetArticle)
		authorized.PUT("/articles/:id", apiArticleHandler.UpdateArticle)
		authorized.DELETE("/articles/:id", apiArticleHandler.DeleteArticle)
	}

	r.Run(":" + todoConfig.Config.Port)

}
