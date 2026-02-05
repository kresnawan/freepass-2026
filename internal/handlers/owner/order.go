package owner

import (
	"canteen/internal/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func GetCanteenOrder(c *gin.Context) {
	cid := c.Param("cid")

	orders, err := sql.SelectOrderByCanteenId(cid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, orders)
}

func GetCanteenOrderById(c *gin.Context) {
	orderIdAny := c.Param("oid")

	oid, err := ulid.Parse(orderIdAny)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	order, _ := sql.SelectOrderById(oid)

	c.JSON(200, order)
}

func UpdateOrderStatus(c *gin.Context) {
	oid := c.Param("oid")
	var status string

	parsedId, err := ulid.Parse(oid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	currentStatus, err := sql.UpdateOrderStatus(parsedId)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	switch currentStatus {
	case 6:
		status = "Waiting payment"
	case 7:
		status = "Cooking"
	case 8:
		status = "Ready"
	case 9:
		status = "Completed"
	default:
		status = ""
	}

	res := fmt.Sprintf("Order status has been updated to %s", status)
	c.String(200, res)
}

func DeleteFeedbackById(c *gin.Context) {
	fid := c.Param("fid")
	err := sql.DeleteFeedback(fid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	res := fmt.Sprintf("Feedback with id %s has deleted", fid)
	c.String(200, res)
}
