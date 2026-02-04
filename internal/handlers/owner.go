package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
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
	res, err := sql.SelectCanteenOwners()
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
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

func AddOwnership(c *gin.Context) {
	cid := c.Param("cid")
	owid := c.Param("owid")

	parsedId, err := ulid.Parse(owid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.InsertOwnership(cid, parsedId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Ownership added")
}
