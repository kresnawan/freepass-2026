package middleware

import (
	"canteen/internal/sql"
	"canteen/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func CheckOwnerOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {

		cid := c.Param("cid")
		mid := c.Param("mid")
		oid := c.Param("oid")
		fid := c.Param("fid")

		accountIdAny, _ := c.Get("account_id")
		accountId, err := utility.AnyToUlid(accountIdAny)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		if cid != "" {
			parsedCanteenId, err := strconv.Atoi(cid)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			isOwned, err := sql.CheckOwnership(accountId, parsedCanteenId)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isOwned {
				c.String(http.StatusForbidden, "You do not own this canteen")
				c.Abort()
				return
			}
		}

		if mid != "" {
			parsedMenuId, err := strconv.Atoi(mid)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			isOwned, err := sql.CheckMenuOwnership(parsedMenuId, accountId)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isOwned {
				c.String(http.StatusForbidden, "You do not own this menu")
				c.Abort()
				return
			}
		}

		if oid != "" {
			parsedOrderId, err := ulid.Parse(oid)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				c.Abort()
				return
			}

			isOwned, err := sql.CheckOrderOwnership(parsedOrderId, accountId)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isOwned {
				c.String(http.StatusForbidden, "You do not own this order")
				c.Abort()
				return
			}
		}

		if fid != "" {
			parsedFeedbackId, err := strconv.Atoi(fid)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			isOwned, err := sql.CheckFeedbackOwnership(parsedFeedbackId, accountId)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isOwned {
				c.String(http.StatusForbidden, "You do not own this feedback")
				c.Abort()
				return
			}
		}
	}
}

func CheckCustomerOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {
		oid := c.Param("oid")

		accountIdAny, _ := c.Get("account_id")
		accountId, err := utility.AnyToUlid(accountIdAny)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		if oid != "" {
			parsedOrderId, err := ulid.Parse(oid)
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				c.Abort()
				return
			}

			isOwned, err := sql.CheckOrderCustomerOwnership(parsedOrderId, accountId)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isOwned {
				c.String(http.StatusForbidden, "You do not own this order")
				c.Abort()
				return
			}
		}
	}
}
