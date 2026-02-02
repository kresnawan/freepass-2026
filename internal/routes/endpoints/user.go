package endpoints

import "github.com/gin-gonic/gin"

func UserEndpointsGroup(rg *gin.RouterGroup) {
	UserEndpoint := rg.Group("/user")
	{
		UserEndpoint.GET("/me/profile")
		UserEndpoint.PUT("/me/profile/edit")
	}
}
