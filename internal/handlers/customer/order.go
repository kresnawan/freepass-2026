package customer

import (
	"canteen/internal/sql"
	"canteen/internal/storage/mariadb"
	"canteen/utility"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func GetMyOrder(c *gin.Context) {
	uid, _ := c.Get("account_id")
	status := c.Query("status")
	page := c.Query("page")
	parsedId, err := utility.AnyToUlid(uid)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	rows, err := sql.SelectMyOrder(parsedId, status, page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", rows))
}

func PlaceOrder(c *gin.Context) {
	uid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(uid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	tx, err := mariadb.Db.Begin()
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	defer tx.Rollback()

	rows, err := sql.SelectMyCart(parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	if len(rows) <= 0 {
		c.JSON(400, api.MakeResponse(0, "Your cart is still empty, none to be ordered", nil))
		c.Abort()
		return
	}

	for _, item := range rows {
		err := sql.CheckMenuStock(item.MenuId, item.Quantity)
		if err != nil {
			c.JSON(400, api.MakeResponse(0, "Order failed, an item in your cart was ran out of stock", nil))
			c.Abort()
			return
		}
	}

	canteenId := rows[0].CanteenId

	oid, err := sql.InsertOrder(parsedId, canteenId, tx)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.InsertOrderItems(oid, rows, tx)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.DeleteMyCartItemsAfterOrder(parsedId, tx)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Order placed, awaiting the payment", nil))
}

func PayOrder(c *gin.Context) {
	oid := c.Param("oid")
	parsedId, err := ulid.Parse(oid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.PayOrder(parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Payment successful", nil))
}

func GetOrderDetails(c *gin.Context) {
	orderIdAny := c.Param("oid")

	orderId, err := ulid.Parse(orderIdAny)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	order, err := sql.GetOrderDetails(orderId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", order))
}

type FeedbackInput struct {
	Description string `json:"description"`
	Star        int    `json:"star"`
}

func AddUserFeedback(c *gin.Context) {
	var reqBody FeedbackInput

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	oid := c.Param("oid")
	parsedOID, err := ulid.Parse(oid)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.AddUserFeedback(parsedOID, reqBody.Description, reqBody.Star)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Feedback added", nil))
}

func GetMyOrderFeedback(c *gin.Context) {
	cusid, _ := c.Get("account_id")
	parsedCusId, err := utility.AnyToUlid(cusid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	feedbacks, err := sql.GetMyOrderFeedback(parsedCusId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", feedbacks))
}
