package middleware

import (
	"canteen/utility/api"
	"canteen/utility/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func proceedToken(c *gin.Context) *jwt.CustomClaims {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Authorization header not found", nil))
		c.Abort()
		return nil
	}

	var authHeaderSplit []string = strings.Split(authHeader, " ")

	if len(authHeaderSplit) != 2 || authHeaderSplit[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Token invalid", nil))
		c.Abort()
		return nil
	}

	var tokenInput string = authHeaderSplit[1]
	token, code, err := jwt.VerifyAccessToken(tokenInput, &jwt.CustomClaims{})

	if err != nil {
		switch code {
		case 242:
			c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Token expired", nil))

		case 243:
			c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Token verification failed", nil))

		case 244:
			c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Token invalid", nil))
		}

		c.Abort()
		return nil
	}

	claims := token.Claims.(*jwt.CustomClaims)
	return claims
}

func CheckIfLoggedIn(c *gin.Context) bool {
	_, err := c.Request.Cookie("refreshToken")

	if err == nil {
		return true
	}

	return false
}

func UserAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		claims := proceedToken(c)
		if claims == nil {
			c.JSON(500, api.MakeResponse(0, "Error when proceeding JWT", nil))
			c.Abort()
			return
		}

		if claims.Role != "customer" && claims.Role != "admin" {
			c.JSON(http.StatusForbidden, api.MakeResponse(0, "Your role do not have permission to access this", nil))
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

		claims := proceedToken(c)
		if claims == nil {
			c.JSON(500, api.MakeResponse(0, "Error when proceeding JWT", nil))
			c.Abort()
			return
		}

		if claims.Role != "owner" && claims.Role != "admin" {
			c.JSON(http.StatusForbidden, api.MakeResponse(0, "Your role do not have permission to access this", nil))
			c.Abort()
			return
		}

		c.Set("account_id", claims.AccountId.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		claims := proceedToken(c)
		if claims == nil {
			c.JSON(500, api.MakeResponse(0, "Error when proceeding JWT", nil))
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, api.MakeResponse(0, "Your role do not have permission to access this", nil))
			c.Abort()
			return
		}

		c.Set("account_id", claims.AccountId.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}
