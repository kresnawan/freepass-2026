package public

import (
	"canteen/internal/models"
	"canteen/internal/sql"
	"canteen/internal/storage/mariadb"
	"canteen/utility/api"
	"canteen/utility/jwt"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func Login(c *gin.Context) {

	type ReqBody struct {
		Cred   string `json:"cred"`
		Passwd string `json:"password"`
	}
	var acc ReqBody

	_ = c.ShouldBindJSON(&acc)

	var passwd string
	var accountId ulid.ULID
	var role string

	query := `
		SELECT 
			account_id, 
			role,
			password
		FROM 
			accounts 
		WHERE 
			(email = ? OR username = ?) AND is_active = 1
	`

	err := mariadb.Db.QueryRow(query, acc.Cred, acc.Cred).Scan(&accountId, &role, &passwd)

	if err != nil {
		if err == mariadb.NoRows {
			c.JSON(404, api.MakeResponse(0, "Email or username not found", nil))
		} else {
			c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		}

		c.Abort()
		return
	}

	match, _, _ := argon2id.CheckHash(acc.Passwd, passwd)

	if !match {
		c.JSON(http.StatusUnauthorized, api.MakeResponse(0, "Wrong password", nil))
		c.Abort()
		return
	}

	accessToken, err := jwt.GenerateAccessToken(accountId, role)
	refreshToken, err := jwt.GenerateRefreshToken(accountId, role)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, "Token generation failed", nil))
		c.Abort()
		return
	}

	c.SetCookie(
		"refreshToken",
		refreshToken,
		60*60*24*7,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(200, api.MakeResponse(1, "", accessToken))
}

func Register(c *gin.Context) {
	var acc models.Account

	_ = c.ShouldBindJSON(&acc)

	exist, err := sql.GetUsernameEmailExistence(acc.Username, acc.Email)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	if !exist {
		c.JSON(400, api.MakeResponse(0, "Username or email already exist", nil))
		c.Abort()
		return
	}

	err = sql.InsertCustomerProfile(acc)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, err.Error(), nil))
		c.Abort()
		return
	}

	c.JSON(200, api.MakeResponse(1, "Register successful", nil))

}

func GetAccessToken(c *gin.Context) {
	refreshToken, err := c.Request.Cookie("refreshToken")

	if err != nil {

		if err == refreshToken.Valid() {
			c.JSON(400, api.MakeResponse(0, "Refresh token invalid", nil))
			c.Abort()
			return
		}

		c.JSON(400, api.MakeResponse(0, "Refresh token not found", nil))
		c.Abort()

		return
	}

	token, _, err := jwt.VerifyRefreshToken(refreshToken.Value, &jwt.CustomClaims{})

	if err != nil {
		c.JSON(400, api.MakeResponse(0, "Refresh token invalid", nil))
		c.Abort()
		return
	}

	claims := token.Claims.(*jwt.CustomClaims)

	newAccessToken, err := jwt.GenerateAccessToken(claims.AccountId, claims.Role)

	if err != nil {
		c.JSON(500, api.MakeResponse(0, "Access token regeneration failed", nil))
		c.Abort()

		return
	}

	c.JSON(200, api.MakeResponse(1, "", newAccessToken))
}

func Logout(c *gin.Context) {
	c.SetCookie("refreshToken", "", -1, "/", "localhost", false, true)
	c.JSON(200, api.MakeResponse(1, "Logged out", nil))
}
