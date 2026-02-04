package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"canteen/utility"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var items []models.CartItem
	uid, _ := c.Get("account_id")

	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = c.ShouldBindJSON(&items)

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

	err = sql.CheckAndInsertToCart(tx, items, parsedId)

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

	c.String(200, "Cart been updated")
}

func GetCart(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	items, err := sql.SelectMyCart(parsedId)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, items)
}

func DeleteCartItems(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.DeleteMyCartItems(parsedId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Cart is now empty")
}
