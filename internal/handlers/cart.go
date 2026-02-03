package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func AddToCart(c *gin.Context) {
	var items []models.CartItem

	err := c.ShouldBindJSON(&items)

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

	for _, item := range items {
		sql.CheckAndInsertToCart(tx, item)
	}

	err = tx.Commit()
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}
}

func GetCart(c *gin.Context) {
	uid, _ := c.Get("account_id")
	s, ok := uid.(string)

	if !ok {
		c.String(500, "Any not string")
		c.Abort()
		return
	}

	parsedId, err := ulid.Parse(s)

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
