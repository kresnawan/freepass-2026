package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func AddMenu(c *gin.Context) {
	cid := c.Param("cid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	ReqBody.CanteenId, _ = strconv.Atoi(cid)
	lastInsertedId, err := sql.InsertMenu(ReqBody)

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
