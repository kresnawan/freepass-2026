package endpoints

import (
	"canteen/internal/handlers/admin"
	"canteen/internal/handlers/owner"
	"canteen/internal/handlers/public"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CanteenEndpointsGroup(rg *gin.RouterGroup) {
	CanteenEndpoint := rg.Group("/canteen")
	{

		CanteenEndpoint.GET("", public.GetCanteen)
		CanteenEndpoint.GET("/:cid/menu", public.GetAllMenuByCanteen)
		CanteenEndpoint.GET("/:cid/feedback", public.GetCanteenFeedbacks)

		OwnerField := CanteenEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		OwnerField.Use(middleware.CheckOwnerOwnership())
		{
			OwnerField.GET("/my", owner.GetOwnedCanteen)
			OwnerField.POST("/:cid/menu", owner.AddMenu)
			OwnerField.GET("/:cid/order", owner.GetCanteenOrder)
			OwnerField.DELETE("/feedback/:fid", owner.DeleteFeedbackById)

		}

		AdminField := CanteenEndpoint.Group("")
		AdminField.Use(middleware.AdminAuth())
		{
			AdminField.POST("", admin.CreateCanteen)
			AdminField.DELETE("/:cid", admin.DeleteCanteen)
			AdminField.GET("/owner", admin.GetAllCanteenOwnership)
			AdminField.POST("/:cid/owner", admin.AddOwnership)
			AdminField.DELETE("/:cid/owner", admin.DeleteOwnership)
			AdminField.GET("/:cid/owner", admin.GetCanteenOwner)

		}
	}

}
