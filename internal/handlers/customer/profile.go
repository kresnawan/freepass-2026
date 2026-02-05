package customer

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedUid, err := utility.AnyToUlid(uid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	res, err := sql.GetCustomerProfile(parsedUid)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func ChangeMyProfile(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedUid, err := utility.AnyToUlid(uid)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	var reqBody models.CustomerProfile

	err = c.ShouldBindJSON(&reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	err = sql.UpdateCustomerProfile(parsedUid, reqBody)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Your profile successfully updated")
}
