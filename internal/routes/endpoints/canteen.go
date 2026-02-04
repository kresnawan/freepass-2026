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
		/* Get all canteen */
		CanteenEndpoint.GET("", handlers.GetCanteen)
		CanteenEndpoint.GET("/menu", handlers.GetAllMenu)
		CanteenEndpoint.GET("/:cid/menu", handlers.GetAllMenuByCanteen)
		CanteenEndpoint.GET("/:cid/feedback", handlers.GetCanteenFeedbacks)

		/* Canteen owner */
		OwnerField := CanteenEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		{
			OwnerField.GET("/my", handlers.GetOwnedCanteen)
			OwnerField.POST("/:cid/menu", handlers.AddMenu)
			OwnerField.PATCH("/:cid/menu/:id", handlers.EditMenu)
			OwnerField.DELETE("/:cid/menu/:id", handlers.DeleteMenuById)
			OwnerField.GET("/:cid/order", handlers.GetCanteenOrder)
			OwnerField.GET("/:cid/order/:oid", handlers.GetCanteenOrderById)
			OwnerField.PATCH("/:cid/order/:oid", handlers.UpdateOrderStatus)
			OwnerField.DELETE("/:cid/feedback/:fid", handlers.DeleteFeedbackById)

		}

		/* Admin */
		AdminField := CanteenEndpoint.Group("")
		AdminField.Use(middleware.AdminAuth())
		{
			AdminField.POST("", handlers.CreateCanteen)

			AdminField.DELETE("/:cid", handlers.DeleteCanteen)

			AdminField.GET("/owner/account", handlers.GetOwnerAccount)
			AdminField.POST("/owner/account", handlers.CreateOwnerAccount)

			AdminField.GET("/owner", handlers.GetAllCanteenOwnership)

			AdminField.POST("/:cid/owner/:owid", handlers.AddOwnership) // Add ownership
			AdminField.GET("/:cid/owner", handlers.GetCanteenOwner)

			AdminField.GET("/:cid/owner/:oid")
			AdminField.PUT("/:cid/owner/:oid")
			AdminField.DELETE("/:cid/owner/:oid")
		}
	}

}
