package owner

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMenu(c *gin.Context) {
	cid := c.Param("cid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	ReqBody.CanteenId, _ = strconv.Atoi(cid)
	err := sql.InsertMenu(ReqBody)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Menu added")
}

func EditMenu(c *gin.Context) {
	mid := c.Param("mid")
	var ReqBody models.Menu

	if err := c.ShouldBindJSON(&ReqBody); err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err := sql.UpdateMenu(mid, ReqBody.MenuName, ReqBody.Price)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Menu updated")
}

func AddMenuStock(c *gin.Context) {
	type RequestBody struct {
		Quantity int `json:"quantity"`
	}

	var reqBody RequestBody
	menuid := c.Param("mid")

	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	menuidInt, err := strconv.Atoi(menuid)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	err = sql.AddStock(menuidInt, reqBody.Quantity)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Stock added")
}

func DeleteMenuById(c *gin.Context) {
	mid := c.Param("mid")

	err := sql.DeleteMenuById(mid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Menu deleted")
}
