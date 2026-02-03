package handlers

import (
	"canteen/internal/handlers/sql"

	"github.com/gin-gonic/gin"
)

func GetCanteen(c *gin.Context) {
	result, err := sql.SelectCanteen()

	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, result)
}

/* Pagination */
func GetMenu(c *gin.Context) {}

func GetMenuByCanteen(c *gin.Context)   {}
func GetMenuById(c *gin.Context)        {}
func GetCanteenFeedback(c *gin.Context) {}
