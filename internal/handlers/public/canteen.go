package public

import (
	"canteen/internal/handlers/sql"

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

func GetAllMenu(c *gin.Context) {
	res, err := sql.GetAllMenu()

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func GetAllMenuByCanteen(c *gin.Context) {
	cid := c.Param("cid")
	res, err := sql.GetAllMenuByCanteen(cid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func GetCanteenFeedbacks(c *gin.Context) {
	cid := c.Param("cid")

	res, err := sql.SelectCanteenFeedback(cid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}
