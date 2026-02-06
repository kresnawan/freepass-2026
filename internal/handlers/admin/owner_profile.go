package admin

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func CreateOwnerAccount(c *gin.Context) {
	var acc models.Account

	err := c.ShouldBindJSON(&acc)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.InsertOwnerProfile(acc)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Owner account created", nil))
}

func GetOwnerAccount(c *gin.Context) {
	res, err := sql.SelectCanteenOwners()
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func UpdateOwnerProfile(c *gin.Context) {
	var reqBody models.OwnerProfileForEdit
	owid := c.Param("owid")
	parsedOwid, err := ulid.Parse(owid)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.UpdateOwnerProfile(parsedOwid, reqBody)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "The owner account has been updated", nil))
}
