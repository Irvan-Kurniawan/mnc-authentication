package controller

import (
	"errors"
	"fmt"
	"mnc-authentication/auth"
	"mnc-authentication/database"
	"mnc-authentication/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankRequest struct {
	Target   string `json:"target"`
	Balance  int    `json:"balance"`
}
func Topup(context *gin.Context) {
	var request BankRequest
	var user entity.Customer

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if request.Balance<=0{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Topup nominal must be larger than 0"})
		context.Abort()
		return
	}
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return
	}
	
	username, err := auth.GetUsername(tokenString, context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	logged := database.Instance.Where("username = ?", username).First(&user)
	if logged.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User not logged in"})
		context.Abort()
		return
	}
	user.Balance+=request.Balance
	update := database.Instance.Save(&user)
	if update.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": update.Error.Error()})
		context.Abort()
		return
	}
	RegisterHistoryTopup(context,user.Username,request.Balance)
	balance := fmt.Sprintf("Customer %v has topped up Rp. %v",user.Username,request.Balance)
	context.JSON(http.StatusOK, gin.H{"message": balance})
}
func Transfer(context *gin.Context) {
	var request BankRequest
	var user entity.Customer
	var target entity.Customer
	
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if request.Balance<=0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Transfer nominal must be larger than 0"})
		context.Abort()
		return
	}
	
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return
	}
	
	username, err := auth.GetUsername(tokenString, context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	logged := database.Instance.Where("username = ?", username).First(&user)
	if logged.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "user not logged in"})
		context.Abort()
		return
	}
	logged2 := database.Instance.Where("username = ?", request.Target).First(&target)
	if logged2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Target does not exist"})
		context.Abort()
		return
	}
	if request.Balance > user.Balance {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Not enough balance"})
		context.Abort()
		return
	}
	if target.Username == user.Username {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Sender must be different than target"})
		context.Abort()
		return
	}
	user.Balance-=request.Balance
	target.Balance+=request.Balance
	update := database.Instance.Save(&user)
	if update.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": update.Error.Error()})
		context.Abort()
		return
	}
	update2 := database.Instance.Save(&target)
	if update2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": update2.Error.Error()})
		context.Abort()
		return
	}
	RegisterHistoryTransfer(context,user.Username,request.Balance)
	balance := fmt.Sprintf("Customer %v has transferred Rp. %v to customer %v",user.Username,request.Balance,target.Username)
	context.JSON(http.StatusOK, gin.H{"message": balance})
}
