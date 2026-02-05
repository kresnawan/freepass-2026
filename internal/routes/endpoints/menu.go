package endpoints

import (
	"canteen/internal/handlers/owner"
	"canteen/internal/handlers/public"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MenuEndpointsGroup(rg *gin.RouterGroup) {
	MenuEndpoint := rg.Group("/menu")
	{

		MenuEndpoint.GET("", public.GetAllMenu)

		OwnerField := MenuEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		OwnerField.Use(middleware.CheckOwnerOwnership())
		{
			OwnerField.PATCH("/:mid", owner.EditMenu)
			OwnerField.PUT("/:mid", owner.AddMenuStock)
			OwnerField.DELETE("/:mid", owner.DeleteMenuById)
		}
	}

}
