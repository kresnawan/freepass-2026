package routes

import (
	"canteen/internal/middleware"
	"canteen/internal/routes/endpoints"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	route := gin.Default()
	route.SetTrustedProxies([]string{"127.0.0.1"})

	route.Use(middleware.CORS())

	v1Endpoint := route.Group("/api/v1")
	{
		endpoints.UserEndpointsGroup(v1Endpoint)
		endpoints.AuthEndpointsGroup(v1Endpoint)
		endpoints.OrderEndpointsGroup(v1Endpoint)
		endpoints.CanteenEndpointsGroup(v1Endpoint)
		endpoints.CartEndpointsGroup(v1Endpoint)
		endpoints.FeedbackEndpointsGroup(v1Endpoint)
		endpoints.MenuEndpointsGroup(v1Endpoint)
		endpoints.OwnerEndpointsGroup(v1Endpoint)
	}

	if err := route.Run(); err != nil {
		fmt.Println("ERROR on running server: ", err)
	}
}
