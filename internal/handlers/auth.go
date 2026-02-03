package handlers

import (
	"canteen/internal/handlers/sql"
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
			email = ? OR username = ?
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

	exist, err := sql.GetUsernameEmailExistence(acc.Username, acc.Email)

	if err != nil {
		c.String(404, "", err.Error())
		c.Abort()
		return
	}

	if !exist {
		c.String(404, "", err.Error())
		c.Abort()
		return
	}

	err = sql.InsertCustomerProfile(acc)

	if err != nil {
		c.String(404, "", err.Error())
		c.Abort()
		return
	}

	c.String(200, "", "Success")

}
