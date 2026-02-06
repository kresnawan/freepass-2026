package customer

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/utility"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedUid, err := utility.AnyToUlid(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	res, err := sql.GetCustomerProfile(parsedUid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func ChangeMyProfile(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedUid, err := utility.AnyToUlid(uid)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	var reqBody models.CustomerProfile

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.UpdateCustomerProfile(parsedUid, reqBody)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Your profile successfully updated", nil))
}
