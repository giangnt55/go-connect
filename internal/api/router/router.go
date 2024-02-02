package router

import "github.com/gin-gonic/gin"

// SetupRouter configures the Gin router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Setup Swagger
	SetupSwagger(r)

	// Define your API routes here
	r.GET("/api/v1/hello", HelloHandler)

	return r
}

// HelloHandler is a simple handler for the "/api/v1/hello" endpoint.
// @Summary Say hello
// @Description Get a friendly greeting
// @Tags greeting
// @Produce json
// @Success 200 {object} string "Hello, World!"
// @Router /api/v1/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
