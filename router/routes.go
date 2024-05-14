package router

import (
	"github.com/becardine/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/oppenings", handler.ListOppeningsHandler)
		v1.POST("/oppening", handler.CreateOppeningHandler)
		v1.GET("/oppening", handler.ShowOppeningHandler)
		v1.PUT("/oppening", handler.UpdateOppeningHandler)
		v1.DELETE("/oppening", handler.DeleteOppeningHandler)
	}
}
