package main

import (
	"mnc-authentication/controller"
	"mnc-authentication/database"
	"mnc-authentication/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root@tcp(localhost:3306)/mnc_db?parseTime=true")
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/customer/register", controller.RegisterCustomer)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
