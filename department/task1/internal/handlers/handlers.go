package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// About implements about endpoint
type About struct {
}

// Get user
func (a *About) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// BlockedUsers from server
type BlockedUsers struct {
	Users map[int]bool
	RWM   sync.RWMutex
}

// User implements user struct
type User struct {
	userID       int
	UsrMinID     int
	BlockedUsers sync.Map
}

// NewUser returns insance of user
func NewUser(usrMinID int) *User {
	return &User{
		UsrMinID: usrMinID,
	}
}

// UnblockUser - delete user from block list
func (u *User) UnblockUser(c *gin.Context) {
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	u.BlockedUsers.Delete(correctID)
	c.JSON(http.StatusOK, gin.H{"Status": userID + " deleted from block list"})
}

// BlockUser - add user to block list
func (u *User) BlockUser(c *gin.Context) {
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	u.BlockedUsers.Store(correctID, true)

	c.JSON(http.StatusOK, gin.H{"Status": userID + " added to block list"})
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
	_, exists := u.BlockedUsers.Load(correctID)

	if exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "user blocked"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}

// GetBlockedUsers returns user
func (u *User) GetBlockedUsers(c *gin.Context) {
	blockedUsers := make(map[string]bool)
	u.BlockedUsers.Range(func(key, value interface{}) bool {
		blockedUsers[fmt.Sprint(key)] = true
		return true
	})

	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"users": blockedUsers})
}
