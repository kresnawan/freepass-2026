package customer

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/internal/storage/mariadb"
	"canteen/utility"
	"canteen/utility/api"
	"canteen/utility/cart"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var items []models.CartItem
	uid, _ := c.Get("account_id")

	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = c.ShouldBindJSON(&items)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	itemsMerged := cart.MergeCartItem(items)

	tx, err := mariadb.Db.Begin()

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	defer tx.Rollback()

	err = sql.CheckAndInsertToCart(tx, itemsMerged, parsedId)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Items added to cart", nil))
}

func GetCart(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	items, err := sql.SelectMyCart(parsedId)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", items))
}

func DeleteCartItems(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.DeleteMyCartItems(parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Your card is now empty", nil))
}
