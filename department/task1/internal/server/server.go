package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// About implements about endpoint
type About struct {
}

func (a *About) get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// User implements user struct
type User struct {
}

func (u *User) getUser(c *gin.Context) {

	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)

	if correctID >= 0 && err == nil {
		c.JSON(http.StatusOK, gin.H{"userId": userID})
		return
	}
	c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
}

// SetupServer - create server and endpoints
func SetupServer() *gin.Engine {
	r := gin.Default()
	var about About
	var userTest User
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.GET("/:userid", userTest.getUser)
			}

			v1.GET("/about", about.get)
		}
	}
	return r
}
