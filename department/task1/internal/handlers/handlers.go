package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// About implements about endpoint
type About struct {
}

// Get user
func (a *About) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// User implements user struct
type User struct {
	userID       int
	UsrMinID     int
	BanndedUsers map[int]bool
}

// UnbanUser - delete user from ban list
func (u *User) UnbanUser(c *gin.Context) {
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	delete(u.BanndedUsers, correctID)
	// fmt.Println(u.BanndedUsers)
	c.JSON(http.StatusOK, gin.H{"Status": userID + " deleted from ban list"})
}

// BanUser - add user to ban list
func (u *User) BanUser(c *gin.Context) {
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	u.BanndedUsers[correctID] = true
	// fmt.Println(u.BanndedUsers)
	c.JSON(http.StatusOK, gin.H{"Status": userID + " added to ban list"})
}

// GetUser returns user
func (u *User) GetUser(c *gin.Context) {
	// Прочитали id юзера и проверили на корректность
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)

	// Если оно не корректно, отпровляем ответ, что все плохо
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}

	if u.BanndedUsers[correctID] {
		c.JSON(http.StatusForbidden, gin.H{"error": "user banned"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}
