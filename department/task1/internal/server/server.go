package server

import (
	"task1/config"
	"task1/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupServer - create server and endpoints
func SetupServer(cfg config.Config) *gin.Engine {
	r := gin.Default()
	var about handlers.About
	var userTest handlers.User
	userTest.UsrMinID = cfg.MinID

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/:userid", userTest.GetUser)
		}

		api.GET("/about", about.Get)

	}
	return r
}
