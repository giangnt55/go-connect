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
	err := setupServer()
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func setupServer() error {
	db, err := database.SqlServerConnect()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the REST API with Gin!",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		var users []entities.User

		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch users",
			})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	err = router.Run(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
