package controller
import (
	"mnc-authentication/database"
	"mnc-authentication/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)
func RegisterCustomer(context *gin.Context) {
	var customer entity.Customer
	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if customer.Username==""{
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Username must not be empty"})
		context.Abort()
		return
	}
	if customer.Password==""{
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Password must not be empty"})
		context.Abort()
		return
	}
	record := database.Instance.Create(&customer)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	RegisterHistoryRegister(context,customer.Username)
	context.JSON(http.StatusCreated, gin.H{"customerId": customer.ID, "username": customer.Username})
}
