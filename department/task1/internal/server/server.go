package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupServer - create server and endpoints
func SetupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})
	r.GET("/api/v1/user/:userid", func(c *gin.Context) {
		userID := c.Param("userid")
		correctID, err := strconv.Atoi(userID)

		if correctID >= 0 && err == nil {
			c.JSON(http.StatusOK, gin.H{"userId": userID})
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		}
	})

	return r
}
