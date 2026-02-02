package endpoints

import "github.com/gin-gonic/gin"

func CanteenEndpointsGroup(rg *gin.RouterGroup) {
	CanteenPubEndpoint := rg.Group("/canteen")
	{
		CanteenPubEndpoint.GET("")
		CanteenPubEndpoint.GET("/menu/:page")
		CanteenPubEndpoint.GET("/:cid/menu")
	}

	CanteenPrivEndpoint := rg.Group("/canteen")
	{
		CanteenPrivEndpoint.GET("/my")

		CanteenPrivEndpoint.POST("/:cid/menu")
		CanteenPrivEndpoint.GET("/:cid/menu/:id")
		CanteenPrivEndpoint.PUT("/:cid/menu/:id/update")
		CanteenPrivEndpoint.DELETE("/:cid/menu/:id/delete")

		CanteenPrivEndpoint.GET("/:cid/order")
		CanteenPrivEndpoint.GET("/:cid/order/:oid")
		CanteenPrivEndpoint.GET("/:cid/order/:oid/feedback")

		CanteenPrivEndpoint.GET("/:cid/feedback/:page")   // All feedbacks
		CanteenPrivEndpoint.DELETE("/:cid/feedback/:fid") // Delete feedback with specific ID

	}

	AdminEndpoint := rg.Group("/canteen")
	{
		AdminEndpoint.POST("")
		AdminEndpoint.GET("/:cid/owner")
		AdminEndpoint.POST("/:cid/owner")

		AdminEndpoint.PUT("/:cid/owner/:oid")
	}
}
