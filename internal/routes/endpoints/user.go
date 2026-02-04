package endpoints

import (
	"canteen/internal/handlers"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserEndpointsGroup(rg *gin.RouterGroup) {
	UserEndpoint := rg.Group("/user")
	UserEndpoint.Use(middleware.UserAuth())
	{
		UserEndpoint.GET("", handlers.GetUsers)
		UserEndpoint.DELETE("/:uid", handlers.DeleteUser)
		UserEndpoint.GET("/profile", handlers.GetMyProfile)
		UserEndpoint.PATCH("/profile", handlers.ChangeMyProfile)
	}
}
