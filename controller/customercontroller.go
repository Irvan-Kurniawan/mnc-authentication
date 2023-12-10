package controller
import (
	"mnc-authentication/database"
	"mnc-authentication/model"
	"net/http"
	"github.com/gin-gonic/gin"
)
func RegisterCustomer(context *gin.Context) {
	var customer model.Customer
	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := customer.HashPassword(customer.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&customer)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"customerId": customer.ID, "username": customer.Username})
}