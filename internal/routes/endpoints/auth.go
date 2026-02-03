package endpoints

import (
	"canteen/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthEndpointsGroup(rg *gin.RouterGroup) {
	AuthEndpoint := rg.Group("/auth")
	{
		AuthEndpoint.POST("/login", handlers.Login)
		AuthEndpoint.POST("/register", handlers.Register)
	}
}
