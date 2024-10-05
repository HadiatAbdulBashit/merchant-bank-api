package router

import (
	"merchant-bank-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create a route group for API
	api := r.Group("/api")
	{
		api.POST("/register", services.Register)
		api.POST("/login", services.Login)
		api.POST("/payment", services.Payment)
		api.POST("/logout", services.Logout)
	}

	return r
}
