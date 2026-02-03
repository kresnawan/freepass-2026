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
	}

	OrderEndpoint := rg.Group("/order")
	OrderEndpoint.Use(middleware.UserAuth())
	{
		OrderEndpoint.GET("", handlers.GetMyOrder)
		OrderEndpoint.POST("", handlers.PlaceOrder)
		OrderEndpoint.POST("/pay")
		OrderEndpoint.POST("/:oid/feedback")
	}
}
