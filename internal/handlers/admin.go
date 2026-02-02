package handlers

import (
	"canteen/internal/handlers/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCanteen(c *gin.Context) {
	type ReqBody struct {
		Name string `json:"name"`
	}

	var requestBody ReqBody
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.String(http.StatusBadRequest, "", "Request body parsing failed")
		c.Abort()
		return
	}

	id, err := sql.AddCanteen(requestBody.Name)

	if err != nil {
		c.String(http.StatusInternalServerError, "", err.Error())
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Success", "inserted_id": id})

}
