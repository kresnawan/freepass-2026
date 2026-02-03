package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCanteen(c *gin.Context) {
	var requestBody models.Canteen
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.String(http.StatusBadRequest, "", "Request body parsing failed")
		c.Abort()
		return
	}

	id, err := sql.AddCanteen(requestBody.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, "", err.Error())
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse("Success", "", 0, id))
}

func DeleteCanteen(c *gin.Context) {
	param := c.Param("cid")
	res, err := sql.DeleteCanteen(param)

	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	if res == 0 {
		c.JSON(404, api.MakeResponse("Canteen not found", "", res))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse("Canteen successfully deleted", "", res))
}

func GetCanteenOwnership(c *gin.Context)    {}
func AddCanteenOwnership(c *gin.Context)    {}
func DeleteCanteenOwnership(c *gin.Context) {}

func CreateCanteenOwnerAccount(c *gin.Context) {
	var acc models.Account

	err := c.ShouldBindJSON(&acc)
	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	err = sql.InsertOwnerProfile(acc)
	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse("Owner account created", ""))
}
func EditCanteenOwnerAccount(c *gin.Context)   {}
func DeleteCanteenOwnerAccount(c *gin.Context) {}

func DisableCustomerAccount(c *gin.Context) {}
