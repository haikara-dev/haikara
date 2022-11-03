package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/haikara-dev/haikara/api"
	todoConfig "github.com/haikara-dev/haikara/config"
	"github.com/haikara-dev/haikara/database"
	"github.com/haikara-dev/haikara/middleware"
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

	apiSiteCategoryHandler := api.SiteCategoryHandler{
		Client: database.Client,
	}

	apiFeedHandler := api.FeedHandler{
		Client: database.Client,
	}

	apiArticleHandler := api.ArticleHandler{
		Client: database.Client,
	}

	apiDashboardHandler := api.DashboardHandler{
		Client: database.Client,
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowOrigins = todoConfig.Config.CORS

	log.Printf("config.AllowOrigins: %v", config.AllowOrigins)

	r.Use(cors.New(config))

	r.GET("/health-check", healthCheck)

	authorized := r.Group("/api")
	authorized.Use(cors.New(config))
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/users/create", apiUserHandler.CreateUserByUUIDAndEmail)
		authorized.GET("/users/current", apiUserHandler.GetCurrentUser)

		authorized.GET("/dashboard", apiDashboardHandler.GetDashboard)
	}

	admin := r.Group("/admin/api")
	admin.Use(cors.New(config))
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware(database.Client))
	{
		admin.GET("/users", apiUserHandler.GetAllUsers)
		admin.POST("/users", apiUserHandler.CreateUser)
		admin.GET("/users/:id", apiUserHandler.GetUser)
		admin.PUT("/users/:id", apiUserHandler.UpdateUser)
		admin.DELETE("/users/:id", apiUserHandler.DeleteUser)
		admin.PATCH("/users/role/:id", apiUserHandler.UpdateUserRole)

		admin.GET("/sites", apiSiteHandler.GetAllSites)
		admin.POST("/sites", apiSiteHandler.CreateSite)
		admin.GET("/sites/:id", apiSiteHandler.GetSite)
		admin.PUT("/sites/:id", apiSiteHandler.UpdateSite)
		admin.DELETE("/sites/:id", apiSiteHandler.DeleteSite)
		admin.PATCH("/sites/active/:id", apiSiteHandler.ActiveSite)
		admin.PATCH("/sites/deActive/:id", apiSiteHandler.DeActiveSite)
		admin.GET("/sites/run-crawling/:id", apiSiteHandler.RunCrawling)
		admin.GET("/sites/dry-run-crawling/:id", apiSiteHandler.DryRunCrawling)
		admin.GET("/sites/get-rss-url/:id", apiSiteHandler.GetRssUrlBySiteId)
		admin.GET("/sites/get-rss-url-by-url", apiSiteHandler.GetRssUrlByUrl)

		admin.GET("/site-categories", apiSiteCategoryHandler.GetSiteCategories)
		admin.POST("/site-categories", apiSiteCategoryHandler.CreateSiteCategory)
		admin.PUT("/site-categories/:id", apiSiteCategoryHandler.UpdateSiteCategory)

		admin.GET("/sites/export", apiSiteHandler.ExportSites)
		admin.POST("/sites/import", apiSiteHandler.ImportSites)

		admin.GET("/feeds", apiFeedHandler.GetAllFeeds)
		admin.GET("/feeds/lite", apiFeedHandler.GetAllFeedsNoneContentsField)
		admin.DELETE("/feeds/:id", apiFeedHandler.DeleteFeed)
		admin.GET("/feeds/parse/:id", apiFeedHandler.ParseFeed)

		admin.GET("/articles", apiArticleHandler.GetAllArticles)
		admin.POST("/articles", apiArticleHandler.CreateArticle)
		admin.GET("/articles/:id", apiArticleHandler.GetArticle)
		admin.PUT("/articles/:id", apiArticleHandler.UpdateArticle)
		admin.DELETE("/articles/:id", apiArticleHandler.DeleteArticle)
		admin.GET("/articles/run-get-ogp-image/:id", apiArticleHandler.RunGetOGPImageOfArticle)

		admin.GET("/dashboard", apiDashboardHandler.GetAdminDashboard)
	}

	r.Run(":" + todoConfig.Config.Port)

}
