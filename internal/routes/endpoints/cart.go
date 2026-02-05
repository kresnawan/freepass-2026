package endpoints

import (
	"canteen/internal/handlers/customer"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CartEndpointsGroup(rg *gin.RouterGroup) {
	CartEndpoint := rg.Group("/cart")
	CartEndpoint.Use(middleware.UserAuth())
	{
		CartEndpoint.GET("", customer.GetCart)
		CartEndpoint.POST("", customer.AddToCart)
		CartEndpoint.DELETE("", customer.DeleteCartItems)
	}
}
