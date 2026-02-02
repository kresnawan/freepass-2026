package routes

import (
	"canteen/internal/middleware"
	"canteen/internal/routes/endpoints"
	"canteen/internal/storage/mariadb"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	route := gin.Default()
	mariadb.DbInit()
	route.SetTrustedProxies([]string{"127.0.0.1"})

	route.Use(middleware.CORS())

	v1Endpoint := route.Group("/api/v1")
	{
		endpoints.UserEndpointsGroup(v1Endpoint)
		endpoints.AuthEndpointsGroup(v1Endpoint)
		endpoints.AdminEndpointsGroup(v1Endpoint)
		endpoints.OrderEndpointsGroup(v1Endpoint)
		endpoints.CanteenEndpointsGroup(v1Endpoint)
	}

	if err := route.Run(); err != nil {
		fmt.Println("ERROR on running server: ", err)
	}
}
