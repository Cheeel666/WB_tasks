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

// BlockedUsers from server
type BlockedUsers struct {
	Users map[int]bool
	RWM   sync.RWMutex
}

// User implements user struct
type User struct {
	userID       int
	UsrMinID     int
	BlockedUsers map[int]bool
	RWM          sync.RWMutex
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
	u.RWM.Lock()
	delete(u.BlockedUsers, correctID)
	u.RWM.Unlock()
	// fmt.Println(u.blockndedUsers)
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
	u.RWM.Lock()
	u.BlockedUsers[correctID] = true
	u.RWM.Unlock()
	// fmt.Println(u.blockndedUsers)
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
	u.RWM.RLock()
	_, exists := u.BlockedUsers[correctID]
	u.RWM.RUnlock()
	if exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "user blocked"})
		return
	}
	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"userId": userID})
}

// GetBlockedUsers returns user
func (u *User) GetBlockedUsers(c *gin.Context) {
	u.RWM.RLock()
	blockedUsers := u.BlockedUsers
	u.RWM.RUnlock()

	// Конечный ответ
	c.JSON(http.StatusOK, gin.H{"users": blockedUsers})
}
