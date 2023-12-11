package middleware

import (
	"mnc-authentication/auth"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
}
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString, context)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		
		context.Next()
	}
}
