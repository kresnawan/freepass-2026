package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	res, err := sql.GetAccounts()

	if err != nil {
		c.String(500, "", err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func DeleteUser(c *gin.Context) {
	uid := c.Param("uid")

	err := sql.DeleteAccount(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse("Success", ""))
}
