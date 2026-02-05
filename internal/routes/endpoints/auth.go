package endpoints

import (
	"canteen/internal/handlers/public"

	"github.com/gin-gonic/gin"
)

func AuthEndpointsGroup(rg *gin.RouterGroup) {
	AuthEndpoint := rg.Group("/auth")
	{
		AuthEndpoint.POST("/login", public.Login)
		AuthEndpoint.POST("/register", public.Register)
		AuthEndpoint.DELETE("/logout", public.Logout)
		AuthEndpoint.GET("/token", public.GetAccessToken)
	}
}
