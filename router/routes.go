package router

import (
	"github.com/becardine/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// initializa handler
	err := handler.Initialize()
	if err != nil {
		panic(err)
	}

	// routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOppeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOppeningHandler)
		v1.PUT("/opening", handler.UpdateOppeningHandler)
		v1.DELETE("/opening", handler.DeleteOppeningHandler)
	}
}
