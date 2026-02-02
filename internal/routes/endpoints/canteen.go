package endpoints

import (
	"canteen/internal/handlers"

	"github.com/gin-gonic/gin"
)

func CanteenEndpointsGroup(rg *gin.RouterGroup) {
	CanteenEndpoint := rg.Group("/canteen")
	{
		/* Public */
		CanteenEndpoint.GET("", handlers.GetCanteen)
		CanteenEndpoint.GET("/menu/:page")
		CanteenEndpoint.GET("/:cid/menu")
		CanteenEndpoint.GET("/:cid/menu/:id")
		CanteenEndpoint.GET("/:cid/feedback/:page")

		/* Canteen owner */
		OwnerField := CanteenEndpoint.Group("")
		{
			OwnerField.GET("/my")
			OwnerField.POST("/:cid/menu")
			OwnerField.PUT("/:cid/menu/:id/update")
			OwnerField.DELETE("/:cid/menu/:id/delete")
			OwnerField.GET("/:cid/order")
			OwnerField.GET("/:cid/order/:oid")
			OwnerField.DELETE("/:cid/feedback/:fid")

		}

		/* Admin */
		AdminField := CanteenEndpoint.Group("")
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
