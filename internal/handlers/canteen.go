package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/utility"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetCanteen(c *gin.Context) {
	result, err := sql.SelectCanteen()

	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, result)
}

func GetOwnedCanteen(c *gin.Context) {
	owid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(owid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	res, err := sql.SelectOwnedCanteen(parsedId)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func GetCanteenOrder(c *gin.Context) {
	cid := c.Param("cid")

	orders, err := sql.SelectOrderByCanteenId(cid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, orders)
}

func GetCanteenOrderById(c *gin.Context) {
	oid := c.Param("oid")

	order, _ := sql.SelectOrderById(oid)

	c.JSON(200, order)
}
