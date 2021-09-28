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
	userID int
}

func (u *User) getUser(c *gin.Context) {
	// Прочитали id юзера и проверили на корректность
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)

	// Если оно не корректно, отпровляем ответ, что все плохо
	if err != nil || correctID < 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}

// SetupServer - create server and endpoints
func SetupServer() *gin.Engine {
	r := gin.Default()
	var about About
	var userTest User
	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/:userid", userTest.getUser)
		}

		api.GET("/about", about.get)

	}
	return r
}
