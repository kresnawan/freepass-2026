package handlers

import (
	"canteen/internal/handlers/sql"
	"canteen/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func GetCanteenFeedbacks(c *gin.Context) {
	cid := c.Param("cid")

	res, err := sql.SelectCanteenFeedback(cid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}

func DeleteFeedbackById(c *gin.Context) {}

type FeedbackInput struct {
	Description string `json:"description"`
	Star        int    `json:"star"`
}

func AddUserFeedback(c *gin.Context) {
	var reqBody FeedbackInput

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	oid := c.Param("oid")
	parsedOID, err := ulid.Parse(oid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.AddUserFeedback(parsedOID, reqBody.Description, reqBody.Star)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Feedback has been added")
}

func GetMyOrderFeedback(c *gin.Context) {
	cusid, _ := c.Get("account_id")
	parsedCusId, err := utility.AnyToUlid(cusid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	feedbacks, err := sql.GetMyOrderFeedback(parsedCusId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, feedbacks)
}
