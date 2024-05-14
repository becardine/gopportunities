package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOppeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
