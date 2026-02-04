package middleware

import (
	"canteen/internal/handlers/sql"
	"canteen/utility/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.String(http.StatusUnauthorized, "ERROR: Authorization header not found")
			c.Abort()
			return
		}

		var authHeaderSplit []string = strings.Split(authHeader, " ")

		if len(authHeaderSplit) != 2 || authHeaderSplit[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			c.Abort()
			return
		}

		var tokenInput string = authHeaderSplit[1]
		token, code, err := jwt.VerifyAccessToken(tokenInput, &jwt.CustomClaims{})

		if err != nil {
			switch code {
			case 242:
				c.String(http.StatusUnauthorized, "ERROR: Token expired")

			case 243:
				c.String(http.StatusUnauthorized, "ERROR: Token verification failed")

			case 244:
				c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			}

			c.Abort()
			return
		}

		claims := token.Claims.(*jwt.CustomClaims)

		if claims.Role != "customer" && claims.Role != "admin" {
			c.String(http.StatusForbidden, "Your role do not have permission to access this")
			c.Abort()
			return
		}

		c.Set("account_id", claims.AccountId.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}

func OwnerAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		cid := c.Param("cid")

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.String(http.StatusUnauthorized, "ERROR: Authorization header not found")
			c.Abort()
			return
		}

		var authHeaderSplit []string = strings.Split(authHeader, " ")

		if len(authHeaderSplit) != 2 || authHeaderSplit[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			c.Abort()
			return
		}

		var tokenInput string = authHeaderSplit[1]
		token, code, err := jwt.VerifyAccessToken(tokenInput, &jwt.CustomClaims{})

		if err != nil {
			switch code {
			case 242:
				c.String(http.StatusUnauthorized, "ERROR: Token expired")

			case 243:
				c.String(http.StatusUnauthorized, "ERROR: Token verification failed")

			case 244:
				c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			}

			c.Abort()
			return
		}

		claims := token.Claims.(*jwt.CustomClaims)

		if claims.Role != "owner" && claims.Role != "admin" {
			c.String(http.StatusForbidden, "Your role do not have permission to access this")
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

			isOwned, err := sql.CheckOwnership(claims.AccountId, parsedCanteenId)
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

		c.Set("account_id", claims.AccountId.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.String(http.StatusUnauthorized, "ERROR: Authorization header not found")
			c.Abort()
			return
		}

		var authHeaderSplit []string = strings.Split(authHeader, " ")

		if len(authHeaderSplit) != 2 || authHeaderSplit[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			c.Abort()
			return
		}

		var tokenInput string = authHeaderSplit[1]
		token, code, err := jwt.VerifyAccessToken(tokenInput, &jwt.CustomClaims{})

		if err != nil {
			switch code {
			case 242:
				c.String(http.StatusUnauthorized, "ERROR: Token expired")

			case 243:
				c.String(http.StatusUnauthorized, "ERROR: Token verification failed")

			case 244:
				c.String(http.StatusUnauthorized, "ERROR: Token invalid")
			}

			c.Abort()
			return
		}

		claims := token.Claims.(*jwt.CustomClaims)

		if claims.Role != "admin" {
			c.String(http.StatusForbidden, "Your role do not have permission to access this")
			c.Abort()
			return
		}

		c.Set("account_id", claims.AccountId.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}
