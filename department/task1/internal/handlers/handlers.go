package handlers

import (
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

// BannedUsers from server
type BannedUsers struct {
	Users map[int]bool
	sync.RWMutex
}

// User implements user struct
type User struct {
	userID   int
	UsrMinID int
	BannedUsers
}

// NewUser returns insance of user
func NewUser(usrMinID int) *User {
	return &User{
		UsrMinID: usrMinID,
	}
}

// UnbanUser - delete user from ban list
func (u *User) UnbanUser(c *gin.Context) {
	userID := c.Param("userid")
	correctID, err := strconv.Atoi(userID)
	if err != nil || correctID < u.UsrMinID {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "user does not exists"})
		return
	}
	u.BannedUsers.RWMutex.Lock()
	delete(u.BannedUsers.Users, correctID)
	u.BannedUsers.RWMutex.Unlock()
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
	u.BannedUsers.RWMutex.Lock()
	u.BannedUsers.Users[correctID] = true
	u.BannedUsers.RWMutex.Unlock()
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

	if u.BannedUsers.Users[correctID] {
		c.JSON(http.StatusForbidden, gin.H{"error": "user banned"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}

// GetBlockedUsers returns user
func (u *User) GetBlockedUsers(c *gin.Context) {
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"users": u.BannedUsers.Users})
}
