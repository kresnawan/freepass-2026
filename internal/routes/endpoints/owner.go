package endpoints

import (
	"canteen/internal/handlers/admin"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OwnerEndpointsGroup(rg *gin.RouterGroup) {
	AdminField := rg.Group("/owner")
	AdminField.Use(middleware.AdminAuth())
	{
		AdminField.GET("", admin.GetOwnerAccount)
		AdminField.POST("", admin.CreateOwnerAccount)
		AdminField.PATCH("", admin.CreateOwnerAccount)
		AdminField.PATCH("/:owid", admin.UpdateOwnerProfile)
	}
}
