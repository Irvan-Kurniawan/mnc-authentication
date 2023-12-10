package controller

import (
	"mnc-authentication/auth"
	"mnc-authentication/database"
	"mnc-authentication/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user entity.Customer
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if username exists and password is correct
	record := database.Instance.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	// check if already logged in
	logged := database.Instance.Where("username = ?", request.Username).Where("is_login=?", 0).First(&user)
	if logged.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Already logged in"})
		context.Abort()
		return
	}
	user.IsLogin = true
	update := database.Instance.Save(&user)
	if update.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": update.Error.Error()})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func Logout(context *gin.Context) {
	var request TokenRequest
	var user entity.Customer
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if already logged out
	logged := database.Instance.Where("username = ?", request.Username).Where("is_login=?", 1).First(&user)
	if logged.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Already logged out"})
		context.Abort()
		return
	}
	user.IsLogin = false
	update := database.Instance.Save(&user)
	if update.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": update.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "Succesfully logged out"})
}
