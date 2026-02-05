package public

import (
	"canteen/internal/handlers/sql"
	"canteen/internal/handlers/sql/accounts"
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
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
			c.String(http.StatusNotFound, "Credential not found")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Abort()
		return
	}

	match, _, _ := argon2id.CheckHash(acc.Passwd, passwd)

	if !match {
		c.String(http.StatusUnauthorized, "Wrong password")
		c.Abort()
		return
	}

	accessToken, err := jwt.GenerateAccessToken(accountId, role)
	refreshToken, err := jwt.GenerateRefreshToken(accountId, role)

	if err != nil {
		c.String(http.StatusInternalServerError, "Token generation failed")
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

	c.Data(200, "", []byte(accessToken))
}

func Register(c *gin.Context) {
	var acc models.Account

	_ = c.ShouldBindJSON(&acc)

	exist, err := accounts.GetUsernameEmailExistence(acc.Username, acc.Email)

	if err != nil {
		c.String(404, err.Error())
		c.Abort()
		return
	}

	if !exist {
		c.String(404, "Username or email have been used")
		c.Abort()
		return
	}

	err = sql.InsertCustomerProfile(acc)

	if err != nil {
		c.String(404, err.Error())
		c.Abort()
		return
	}

	c.String(200, "", "Success")

}

func GetAccessToken(c *gin.Context) {
	refreshToken, err := c.Request.Cookie("refreshToken")

	if err != nil {

		if err == refreshToken.Valid() {
			c.Data(http.StatusBadRequest, "", []byte("ERROR: Refresh token invalid"))
			c.Abort()

			return
		}

		c.Data(http.StatusBadRequest, "", []byte("ERROR: Refresh token not found"))
		c.Abort()

		return
	}

	token, _, err := jwt.VerifyRefreshToken(refreshToken.Value, &jwt.CustomClaims{})

	if err != nil {
		c.Data(http.StatusBadRequest, "", []byte("ERROR: Refresh token invalid"))
		c.Abort()
		return
	}

	claims := token.Claims.(*jwt.CustomClaims)

	newAccessToken, err := jwt.GenerateAccessToken(claims.AccountId, claims.Role)

	if err != nil {
		c.Data(500, "", []byte("ERROR: Access token regeneration failed"))
		c.Abort()

		return
	}

	c.Data(200, "", []byte(newAccessToken))
}

func Logout(c *gin.Context) {
	c.SetCookie("refreshToken", "", -1, "/", "localhost", false, true)
	c.String(200, "Logged out")
}
