package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func CreateOwnerAccount(c *gin.Context) {
	var acc models.Account

	err := c.ShouldBindJSON(&acc)
	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	err = sql.InsertOwnerProfile(acc)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse("Owner account created", ""))
}

func GetOwnerAccount(c *gin.Context) {

}

func GetAllCanteenOwnership(c *gin.Context) {
	res, err := sql.SelectAllCanteenOwnership()

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func GetCanteenOwner(c *gin.Context) {
	cid := c.Param("cid")
	res, err := sql.SelectCanteenOwner(cid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}
