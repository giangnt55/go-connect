package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initializing REST API...")

	// Create a default Gin router
	router := gin.Default()

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the REST API with Gin!",
		})
	})

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
