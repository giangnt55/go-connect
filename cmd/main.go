package main

import (
	"fmt"
	"go-connect/pkg/domain/entities"
	"go-connect/pkg/infrastructure/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initializing REST API...")

	db, err := database.SqlServerConnect()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	// Create a default Gin router
	router := gin.Default()

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the REST API with Gin!",
		})
	})

	// Example route that interacts with the database
	router.GET("/users", func(c *gin.Context) {
		var users []entities.User

		// Fetch users from the database using GORM
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	// Run the server on port 8080
	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
