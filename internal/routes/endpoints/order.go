package endpoints

import (
	"canteen/internal/handlers/customer"
	"canteen/internal/handlers/owner"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderEndpointsGroup(rg *gin.RouterGroup) {
	OrderEndpoint := rg.Group("/order")
	{
		CustomerField := OrderEndpoint.Group("")
		CustomerField.Use(middleware.UserAuth())
		CustomerField.Use(middleware.CheckCustomerOwnership())
		{
			CustomerField.GET("", customer.GetMyOrder)
			CustomerField.POST("", customer.PlaceOrder)
			CustomerField.GET("/:oid", customer.GetOrderDetails)
			CustomerField.POST("/:oid/pay", customer.PayOrder)
			CustomerField.POST("/:oid/feedback", customer.AddUserFeedback)
		}

		OwnerField := OrderEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		OwnerField.Use(middleware.CheckOwnerOwnership())
		{
			OwnerField.GET("/:oid", owner.GetCanteenOrderById)
			OwnerField.PATCH("/:oid", owner.UpdateOrderStatus)
		}
	}
}
