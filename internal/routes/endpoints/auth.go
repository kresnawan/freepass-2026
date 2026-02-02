package endpoints

import "github.com/gin-gonic/gin"

func AuthEndpointsGroup(rg *gin.RouterGroup) {
	AuthEndpoint := rg.Group("/auth")
	{
		AuthEndpoint.POST("/login")
		AuthEndpoint.POST("/register")
	}
}
