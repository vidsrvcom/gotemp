// Package main is the entry point of the application.
// This file provides a convenient way to run the application from the project root.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vidsrvcom/gotemp/internal/config"
	"github.com/vidsrvcom/gotemp/internal/handler"
	"github.com/vidsrvcom/gotemp/internal/repository"
	"github.com/vidsrvcom/gotemp/internal/server"
	"github.com/vidsrvcom/gotemp/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger := log.New(os.Stdout, "[GOTEMP] ", log.LstdFlags|log.Lshortfile)

	// Initialize repository layer
	userRepo := repository.NewUserRepository()

	// Initialize service layer
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	homeHandler := handler.NewHomeHandler()

	// Create and configure the server
	srv := server.New(cfg, logger, homeHandler, userHandler)

	// Start server in a goroutine
	go func() {
		logger.Printf("Starting server on %s", cfg.ServerAddress())
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exited gracefully")
}
