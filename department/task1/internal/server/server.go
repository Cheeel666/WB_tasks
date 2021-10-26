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
	var about handlers.About
	var userTest handlers.User
	userTest.UsrMinID = cfg.MinID
	userTest.BanndedUsers = utils.AddBanned(cfg.BannedUsrs)
	// fmt.Println("BannedUsers:", userTest.BanndedUsers)

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/:userid", userTest.GetUser)
			user.DELETE("/unban/:userid", userTest.UnbanUser)
			user.PUT("/ban/:userid", userTest.BanUser)
		}

		api.GET("/about", about.Get)

	}
	return r
}
