package routes

import (
	"canteen/internal/storage/mariadb"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRoute() {

	route := gin.Default()
	mariadb.DbInit()
	route.SetTrustedProxies([]string{"127.0.0.1"})

	route.GET("", func(c *gin.Context) {
		c.String(200, "Hello world!")
		c.Abort()
	})

	if err := route.Run(); err != nil {
		fmt.Println("ERROR on running server: ", err)
	}
}
