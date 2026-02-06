package routes

import (
	"canteen/internal/middleware"
	"canteen/internal/routes/endpoints"
	"log"
	"os"

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

	log.Printf("Starting BCC Canteen on :%s...", os.Getenv("PORT"))
	if err := route.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
