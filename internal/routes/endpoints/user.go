package endpoints

import (
	"canteen/internal/handlers"

	"github.com/gin-gonic/gin"
)

func UserEndpointsGroup(rg *gin.RouterGroup) {
	UserEndpoint := rg.Group("/user")
	{
		UserEndpoint.GET("", handlers.GetUsers)
		UserEndpoint.DELETE("/:uid", handlers.DeleteUser)
		UserEndpoint.GET("/me/profile")
		UserEndpoint.PUT("/me/profile/edit")
	}
}
