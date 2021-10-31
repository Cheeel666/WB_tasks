package server

import (
	"task1/config"
	"task1/internal/handlers"
	"task1/internal/utils"

	"github.com/gin-gonic/gin"
)

// SetupServer - create server and endpoints
func SetupServer(cfg config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(
		CORSMiddleware(),
	)

	var about handlers.About
	userTest := handlers.NewUser(cfg.MinID)
	userTest.BannedUsers.Users = utils.AddBanned(cfg.BannedUsrs)

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/:userid", userTest.GetUser)
			user.GET("/blocked-users", userTest.GetBlockedUsers)
			user.DELETE("/:userid/block", userTest.UnbanUser)
			user.PUT("/:userid/block", userTest.BanUser)
		}

		api.GET("/about", about.Get)

	}
	return r
}

// CORSMiddleware for cross origin requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
