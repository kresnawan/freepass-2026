package owner

import (
	"canteen/internal/sql"
	"canteen/utility/api"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func GetCanteenOrder(c *gin.Context) {
	cid := c.Param("cid")
	status := c.Query("status")
	page := c.Query("page")

	orders, err := sql.SelectOrderByCanteenId(cid, status, page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", orders))
}

func GetCanteenOrderById(c *gin.Context) {
	orderIdAny := c.Param("oid")

	oid, err := ulid.Parse(orderIdAny)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	order, _ := sql.SelectOrderById(oid)

	c.JSON(200, api.MakeResponse(1, "", order))
}

func UpdateOrderStatus(c *gin.Context) {
	oid := c.Param("oid")
	var status string

	parsedId, err := ulid.Parse(oid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	currentStatus, err := sql.UpdateOrderStatus(parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
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
	c.JSON(200, api.MakeResponse(1, res, nil))
}

func DeleteFeedbackById(c *gin.Context) {
	fid := c.Param("fid")
	err := sql.DeleteFeedback(fid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	res := fmt.Sprintf("Feedback with id %s has deleted", fid)
	c.JSON(200, api.MakeResponse(1, "", res))
}
