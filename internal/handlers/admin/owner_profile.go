package admin

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"
	"net/http"

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

func UpdateOwnerProfile(c *gin.Context) {
	var reqBody models.OwnerProfileForEdit
	owid := c.Param("owid")
	parsedOwid, err := ulid.Parse(owid)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	err = sql.UpdateOwnerProfile(parsedOwid, reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.String(200, "The owner account has been updated")
}
