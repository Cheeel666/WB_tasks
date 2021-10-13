package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// About implements about endpoint
type About struct {
}

func (a *About) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// User implements user struct
type User struct {
	userID   int
	UsrMinID int
}

func (u *User) GetUser(c *gin.Context) {
	// Прочитали id юзера и проверили на корректность
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)

	// Если оно не корректно, отпровляем ответ, что все плохо
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}
