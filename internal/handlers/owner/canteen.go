package owner

import (
	"canteen/internal/sql"
	"canteen/utility"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
)

func GetOwnedCanteen(c *gin.Context) {
	owid, _ := c.Get("account_id")
	parsedId, err := utility.AnyToUlid(owid)

	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	res, err := sql.SelectOwnedCanteen(parsedId)

	if err != nil {
		c.JSON(500, api.MakeResponse("Error", err.Error()))
		c.Abort()
		return
	}

	c.JSON(200, res)
}
