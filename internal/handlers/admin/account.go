package admin

import (
	"canteen/internal/sql"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func DeactiveAccount(c *gin.Context) {
	uid := c.Param("uid")
	parsedUid, err := ulid.Parse(uid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	err = sql.DeactiveAccount(parsedUid)
	if err != nil {
		c.String(500, err.Error())
		c.Abort()
		return
	}

	c.String(200, "Account deactived and unusable until activated again")
}

func GetAllAccounts(c *gin.Context) {
	res, err := sql.GetAccounts()

	if err != nil {
		c.String(500, "", err.Error())
		c.Abort()
		return
	}

	c.JSON(200, res)
}
