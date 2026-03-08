package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vashkevich/blockchain/internal/app"
	"github.com/vashkevich/blockchain/internal/config"
	"github.com/vashkevich/blockchain/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := logger.New(cfg.LogLevel)

	logger.Info("Starting Crypto Payment Gateway...")
	logger.Info("Server will listen",
		zap.String("host", cfg.ServerHost),
		zap.String("port", cfg.ServerPort),
	)

	// Create application
	application, err := app.New(cfg, logger)
	if err != nil {
		logger.Error("Failed to create application", zap.Error(err))
		os.Exit(1)
	}

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run application
	go func() {
		if err := application.Run(ctx); err != nil {
			logger.Error("Application error", zap.Error(err))
			cancel()
		}
	}()

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	logger.Info("Shutting down gracefully...")

	if err := application.Shutdown(ctx); err != nil {
		logger.Error("Error during shutdown", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Application stopped")
}
