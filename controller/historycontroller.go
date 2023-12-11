package controller

import (
	"mnc-authentication/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
type History struct {
	Username string
	Action string
	Nominal  int
	ActionDate time.Time
}
func RegisterHistoryLogin(context *gin.Context, username string) {
	
	history := History{Username: username, Action: "Login", ActionDate: time.Now()}

	result := database.Instance.Create(&history)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
}
func RegisterHistoryLogout(context *gin.Context, username string) {
	
	history := History{Username: username, Action: "Logout", ActionDate: time.Now()}

	result := database.Instance.Create(&history)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
}
func RegisterHistoryTopup(context *gin.Context, username string, nominal int) {
	
	history := History{Username: username, Action: "Topup", Nominal: nominal, ActionDate: time.Now()}

	result := database.Instance.Create(&history)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
}
func RegisterHistoryTransfer(context *gin.Context, username string, nominal int) {
	
	history := History{Username: username, Action: "Transfer", Nominal: nominal, ActionDate: time.Now()}

	result := database.Instance.Create(&history)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
}
func RegisterHistoryRegister(context *gin.Context, username string) {
	
	history := History{Username: username, Action: "Register", ActionDate: time.Now()}

	result := database.Instance.Create(&history)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
}
