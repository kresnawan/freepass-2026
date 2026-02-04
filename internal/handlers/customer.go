package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/models"
	"canteen/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

// func MakeOrder(c *gin.Context) {}
// func PayOrder(c *gin.Context)  {}

// func GetMyOrder(c *gin.Context)   {}
// func GetOrderById(c *gin.Context) {}
// func GiveFeedback(c *gin.Context) {}

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

	err = sql.UpdateUserProfile(parsedUid, reqBody)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Your profile successfully updated")
}
