package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetOwnedCanteen(c *gin.Context) {
	owid := c.Param("owid")
	res, err := sql.SelectOwnedCanteen(owid)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, res)
}
func AddMenu(c *gin.Context) {
	cid := c.Param("cid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	lastInsertedId, err := sql.InsertMenu(cid, ReqBody.MenuName, ReqBody.Price)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(500, api.MakeResponse("Menu successfully added", "", 0, lastInsertedId))
}
func EditMenu(c *gin.Context) {
	mid := c.Param("mid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	rowsAffected, err := sql.UpdateMenu(mid, ReqBody.MenuName, ReqBody.Price)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(500, api.MakeResponse("Menu successfully patched", "", rowsAffected))
}
func DeleteMenuById(c *gin.Context) {
	mid := c.Param("mid")

	rowsAffected, err := sql.DeleteMenuById(mid)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
	}

	c.JSON(200, api.MakeResponse("Menu successfully deleted", "", rowsAffected))
}
func GetCanteenOrder(c *gin.Context) {
	cid := c.Param("cid")

	orders, _ := sql.SelectOrderByCanteenId(cid)

	c.JSON(200, orders)
}
func GetCanteenOrderById(c *gin.Context) {
	oid := c.Param("oid")

	order, _ := sql.SelectOrderById(oid)

	c.JSON(200, order)
}
func DeleteFeedbackById(c *gin.Context) {}
