package endpoints

import (
	"canteen/internal/handlers/admin"
	"canteen/internal/handlers/customer"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserEndpointsGroup(rg *gin.RouterGroup) {
	UserEndpoint := rg.Group("/user")
	{
		CustomerField := UserEndpoint.Group("")
		CustomerField.Use(middleware.UserAuth())
		{
			CustomerField.GET("/profile", customer.GetMyProfile)
			CustomerField.PATCH("/profile", customer.ChangeMyProfile)
		}

		AdminField := rg.Group("/user")
		AdminField.Use(middleware.AdminAuth())
		{
			AdminField.PATCH("/:uid")
			AdminField.GET("", admin.GetAllAccounts)
			AdminField.DELETE("/:uid", admin.DeactiveAccount)
		}
	}

}
