package admin

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func CreateCanteen(c *gin.Context) {
	var requestBody models.Canteen
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	_, err = sql.AddCanteen(requestBody.Name)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Canteen created", nil))
}

func DeleteCanteen(c *gin.Context) {
	param := c.Param("cid")
	res, err := sql.DeleteCanteen(param)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	if res == 0 {
		c.JSON(404, api.MakeResponse(0, "Canteen not found", nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Canteen deactivated", nil))
}

func ReactivateCanteen(c *gin.Context) {
	param := c.Param("cid")
	err := sql.ReactivateCanteen(param)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Canteen reactivated", nil))
}

func GetAllCanteenOwnership(c *gin.Context) {
	res, err := sql.SelectAllCanteenOwnership()

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func GetCanteenOwner(c *gin.Context) {
	cid := c.Param("cid")
	res, err := sql.SelectCanteenOwner(cid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}

func AddOwnership(c *gin.Context) {
	cid := c.Param("cid")

	var reqBody struct {
		OwnerId string `json:"owner_id"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	parsedId, err := ulid.Parse(reqBody.OwnerId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.InsertOwnership(cid, parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Canteen ownership added", nil))
}

func DeleteOwnership(c *gin.Context) {
	cid := c.Param("cid")

	var reqBody struct {
		OwnerId string `json:"owner_id"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	parsedId, err := ulid.Parse(reqBody.OwnerId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.DeleteOwnership(cid, parsedId)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Canteen ownership removed", nil))
}
