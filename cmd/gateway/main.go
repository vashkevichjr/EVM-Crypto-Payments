package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vashkevich/blockchain/internal/app"
	"github.com/vashkevich/blockchain/internal/config"
	"github.com/vashkevich/blockchain/internal/storage"
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

	// Create database connection pool
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancel()
	db, err := storage.NewPostgresPool(dbCtx, cfg.DSN())
	if err != nil {
		logger.Error("Failed to create database pool", zap.Error(err))
		os.Exit(1)
	}
	defer db.Close()

	// Create application
	application, err := app.New(cfg, logger, db)
	if err != nil {
		logger.Error("Failed to create application", zap.Error(err))
		os.Exit(1)
	}

	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	// Graceful shutdown
	signalCtx, signalStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer signalStop()

	// Run application
	go func() {
		if err := application.Run(appCtx); err != nil {
			logger.Error("Application error", zap.Error(err))
			appCancel()
		}
	}()

	select {
	case <-signalCtx.Done():
		logger.Info("Received OS signal, shutting down gracefully...")
	case <-appCtx.Done():
		logger.Info("Application crashed, initiating shutdown...")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := application.Shutdown(shutdownCtx); err != nil {
		logger.Error("Error during shutdown", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Application stopped")
	os.Exit(0)
}
