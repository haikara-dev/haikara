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

func authorizedApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func StartWebserver() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowOrigins = todoConfig.Config.CORS

	r.Use(cors.New(config))

	r.GET("/health-check", healthCheck)

	apiUserHandler := api.UserHandler{
		Client: database.Client,
	}

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/user/create", apiUserHandler.CreateUserByUUID)
		authorized.GET("/authorized", authorizedApi)
	}

	//r.GET("/user", apiUserHandler.GetAllUsers)
	//r.POST("/user", apiUserHandler.CreateUser)
	//r.GET("/user/:id", apiUserHandler.GetUser)
	//r.PUT("/user/:id", apiUserHandler.UpdateUser)
	//r.DELETE("/user/:id", apiUserHandler.DeleteUser)

	r.Run(":" + todoConfig.Config.Port)

}
