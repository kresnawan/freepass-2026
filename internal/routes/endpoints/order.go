package endpoints

import (
	"canteen/internal/handlers"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderEndpointsGroup(rg *gin.RouterGroup) {
	/* Only for logged-in user */
	CartEndpoint := rg.Group("/cart")
	CartEndpoint.Use(middleware.UserAuth())
	{
		CartEndpoint.GET("", handlers.GetCart)
		CartEndpoint.POST("", handlers.AddToCart)
		CartEndpoint.DELETE("", handlers.DeleteCartItems)
	}

	OrderEndpoint := rg.Group("/order")
	OrderEndpoint.Use(middleware.UserAuth())
	{
		OrderEndpoint.GET("", handlers.GetMyOrder)
		OrderEndpoint.POST("", handlers.PlaceOrder)
		OrderEndpoint.POST("/:oid/pay", handlers.PayOrder)
		OrderEndpoint.POST("/:oid/feedback", handlers.AddUserFeedback)
		OrderEndpoint.GET("/:oid/feedback", handlers.GetMyOrderFeedback)
	}
}
