package main

import (
	"go-connect/internal/api/router"
	"go-connect/internal/api/server"
)

func main() {
	// Initialize the Gin router
	r := router.SetupRouter()

	// Initialize the HTTP server
	s := server.NewServer(r, 8080)

	// Start the server
	if err := s.Start(); err != nil {
		panic(err)
	}
}
