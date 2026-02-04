package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/storage/mariadb"
	"canteen/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

/*

OrderEndpoint.GET("")
OrderEndpoint.GET("/:oid")
OrderEndpoint.POST("/place")
OrderEndpoint.POST("/:oid/pay")
OrderEndpoint.POST("/:oid/feedback")

*/

func GetMyOrder(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	rows, err := sql.SelectMyOrder(parsedId)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, rows)
}

func PlaceOrder(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}
	tx, err := mariadb.Db.Begin()

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	defer tx.Rollback()

	rows, err := sql.SelectMyCart(parsedId)

	if len(rows) <= 0 {
		c.String(http.StatusBadRequest, "Your cart is still empty, none to be ordered")
		c.Abort()
		return
	}

	canteenId := rows[0].CanteenId

	oid, err := sql.InsertOrder(parsedId, canteenId, tx)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.InsertOrderItems(oid, rows, tx)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.DeleteMyCartItemsAfterOrder(parsedId, tx)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = tx.Commit()
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Order placed")
}

func UpdateOrderStatus(c *gin.Context) {
	oid := c.Param("oid")

	parsedId, err := ulid.Parse(oid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.UpdateOrderStatus(parsedId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Order status has been updated")
}

func PayOrder(c *gin.Context) {
	oid := c.Param("oid")
	parsedId, err := ulid.Parse(oid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.PayOrder(parsedId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Payment successful")
}
