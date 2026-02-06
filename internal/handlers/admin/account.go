package admin

import (
	"canteen/internal/sql"
	"canteen/utility/api"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func DeactiveAccount(c *gin.Context) {
	uid := c.Param("uid")
	parsedUid, err := ulid.Parse(uid)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.DeactiveAccount(parsedUid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Account deactivated and unusable until activated again", nil))
}

func ReactivateAccount(c *gin.Context) {
	uid := c.Param("uid")
	parsedUid, err := ulid.Parse(uid)
	if err != nil {
		c.JSON(400, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	err = sql.ReactivateAccount(parsedUid)
	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Account reactivated", nil))
}

func GetAllAccounts(c *gin.Context) {
	page := c.Query("page")
	res, err := sql.GetAccounts(page)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "", res))
}
