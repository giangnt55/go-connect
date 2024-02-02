// server.go

package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server.
type Server struct {
	httpServer *http.Server
}

// NewServer creates a new HTTP server with the provided router.
func NewServer(handler http.Handler, port int) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		},
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	if ginEngine, ok := s.httpServer.Handler.(*gin.Engine); ok {
		SetupCORS(ginEngine)
	} else {
		fmt.Println("Unable to setup CORS: not a *gin.Engine")
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	fmt.Printf("Server started on %s\n", s.httpServer.Addr)

	// Listen for interrupt signals to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	fmt.Println("Server gracefully stopped")

	return nil
}
