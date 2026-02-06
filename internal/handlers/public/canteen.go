package public

import (
	"canteen/internal/sql"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetCanteen(c *gin.Context) {
	page := c.Query("page")
	result, err := sql.SelectCanteen(page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", result))
}

func GetAllMenu(c *gin.Context) {
	page := c.Query("page")
	res, err := sql.GetAllMenu(page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func GetAllMenuByCanteen(c *gin.Context) {
	cid := c.Param("cid")
	page := c.Query("page")

	res, err := sql.GetAllMenuByCanteen(cid, page)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func GetCanteenFeedbacks(c *gin.Context) {
	cid := c.Param("cid")
	page := c.Query("page")

	res, err := sql.SelectCanteenFeedback(cid, page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}
