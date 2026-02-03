package handlers

import (
	"canteen/internal/handlers/sql"

	"github.com/gin-gonic/gin"
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
