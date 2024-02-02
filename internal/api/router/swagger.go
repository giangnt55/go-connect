package router

import (
	_ "go-connect/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupSwagger configures Swagger documentation.
func SetupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up the path for swagger.json
	router.StaticFile("/swagger.json", "./swagger.json")
}
