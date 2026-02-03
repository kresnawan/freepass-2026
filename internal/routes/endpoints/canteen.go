package endpoints

import (
	"canteen/internal/handlers"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CanteenEndpointsGroup(rg *gin.RouterGroup) {
	CanteenEndpoint := rg.Group("/canteen")
	{
		/* Public */
		CanteenEndpoint.GET("", handlers.GetCanteen)
		CanteenEndpoint.GET("/menu/:page")
		CanteenEndpoint.GET("/:cid/menu")
		CanteenEndpoint.GET("/:cid/feedback/:page")

		/* Canteen owner */
		OwnerField := CanteenEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		{
			OwnerField.GET("/my")
			OwnerField.POST("/:cid/menu")
			OwnerField.PATCH("/:cid/menu/:id")
			OwnerField.DELETE("/:cid/menu/:id")
			OwnerField.GET("/:cid/order")
			OwnerField.GET("/:cid/order/:oid")
			OwnerField.PATCH("/:cid/order/:oid")
			OwnerField.DELETE("/:cid/feedback/:fid")

		}

		/* Admin */
		AdminField := CanteenEndpoint.Group("")
		AdminField.Use(middleware.AdminAuth())
		{
			AdminField.POST("", handlers.CreateCanteen)
			AdminField.DELETE("/:cid", handlers.DeleteCanteen)

			AdminField.POST("/:cid/owner")
			AdminField.GET("/:cid/owner")

			AdminField.GET("/:cid/owner/:oid")
			AdminField.PUT("/:cid/owner/:oid")
			AdminField.DELETE("/:cid/owner/:oid")
		}
	}

}
