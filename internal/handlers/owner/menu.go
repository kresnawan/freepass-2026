package owner

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/utility/api"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMenu(c *gin.Context) {
	cid := c.Param("cid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	ReqBody.CanteenId, _ = strconv.Atoi(cid)
	err := sql.InsertMenu(ReqBody)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Menu added", nil))
}

func EditMenu(c *gin.Context) {
	mid := c.Param("mid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err := sql.UpdateMenu(mid, ReqBody.MenuName, ReqBody.Price)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Menu updated", nil))
}

func AddMenuStock(c *gin.Context) {
	type RequestBody struct {
		Quantity int `json:"quantity"`
	}

	var reqBody RequestBody
	menuid := c.Param("mid")

	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	menuidInt, err := strconv.Atoi(menuid)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.AddStock(menuidInt, reqBody.Quantity)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Stock added", nil))
}

func DeleteMenuById(c *gin.Context) {
	mid := c.Param("mid")

	err := sql.DeleteMenuById(mid)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Menu removed", nil))
}
