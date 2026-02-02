package endpoints

import "github.com/gin-gonic/gin"

func OrderEndpointsGroup(rg *gin.RouterGroup) {
	OrderEndpoint := rg.Group("/order")
	{
		OrderEndpoint.POST("/place")
		OrderEndpoint.POST("/:oid/pay")
		OrderEndpoint.GET("/:oid")
		OrderEndpoint.POST("/:oid/feedback")
	}
}
