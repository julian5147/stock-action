package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"stockapi/internal/application"
	"stockapi/internal/infrastructure/api"
	"stockapi/internal/infrastructure/config"
	"stockapi/internal/infrastructure/external/stockapi"
	"stockapi/internal/infrastructure/logging"
	"stockapi/internal/infrastructure/persistence/cockroach"
)

func main() {
	// Create a cancelable context for graceful shutdown handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configure signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}

	// Initialize logger
	logger := logging.NewStockLogger()

	// Initialize repositories and clients with logger
	stockRepo, err := cockroach.NewStockRepository(ctx, cfg.DatabaseURL, logger)
	if err != nil {
		log.Fatalf("error initializing stock repository: %v", err)
	}
	apiClient := stockapi.NewStockAPIClient(cfg.ExternalAPIURL, cfg.AuthToken, logger)

	// Initialize application with WebSocket handler
	app := application.NewStockApplication(stockRepo, apiClient, logger)

	// Initialize and run server
	server := api.NewServer(cfg, app)

	// Run server in a goroutine
	go func() {
		if err := server.Run(); err != nil {
			log.Printf("error in server: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	<-sigChan
	log.Println("starting graceful shutdown...")

	// Create context with timeout for shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	// Perform cleanup and shutdown
	if err := stockRepo.Close(shutdownCtx); err != nil {
		log.Printf("error during shutdown: %v", err)
	}

	log.Println("server stopped correctly")
}
