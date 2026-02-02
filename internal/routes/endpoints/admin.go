package endpoints

import "github.com/gin-gonic/gin"

func AdminEndpointsGroup(rg *gin.RouterGroup) {

	AdminEndpoint := rg.Group("/admin")
	{
		AdminEndpoint.POST("/canteen")
		AdminEndpoint.POST("/canteen/:cid/owner")
		AdminEndpoint.GET("/canteen/:cid/owner")
		AdminEndpoint.GET("/canteen/:cid/owner/:oid")
		AdminEndpoint.PUT("/canteen/:cid/owner/:oid")
		AdminEndpoint.DELETE("/canteen/:cid/owner/:oid")
		AdminEndpoint.DELETE("/user/:uid")
	}
}
