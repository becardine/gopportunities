package router

import (
	"github.com/becardine/gopportunities/handler"
	"github.com/gin-gonic/gin"

	docs "github.com/becardine/gopportunities/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// initializa handler
	err := handler.Initialize()
	if err != nil {
		panic(err)
	}

	basePath := "/api/v1"

	// swagger docs
	docs.SwaggerInfo.Title = "Gopportunities API"
	docs.SwaggerInfo.Description = "This is a sample server Gopportunities server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = basePath

	// routes
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	// initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
