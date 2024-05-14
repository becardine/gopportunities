package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOppeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
