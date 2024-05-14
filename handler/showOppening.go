package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOppeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
