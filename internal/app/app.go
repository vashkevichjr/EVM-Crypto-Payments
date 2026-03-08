package app

import (
	"context"
	"fmt"

	"github.com/vashkevich/blockchain/internal/config"
	"github.com/vashkevich/blockchain/pkg/logger"
)

// App represents the main application
type App struct {
	config *config.Config
	logger *logger.Logger
}

// New creates a new application
func New(cfg *config.Config, log *logger.Logger) (*App, error) {
	app := &App{
		config: cfg,
		logger: log,
	}

	// Database, routers, and workers initialization will be here
	// For now, just return the basic structure

	return app, nil
}

// Run starts the application
func (a *App) Run(ctx context.Context) error {
	a.logger.Info("Application started")

	// HTTP server and workers startup will be here
	// For now, just wait for context

	<-ctx.Done()
	return ctx.Err()
}

// Shutdown gracefully stops the application
func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Shutting down application...")

	// Database connections closing and workers shutdown will be here
	// For now, just log

	return nil
}

// GetAddress returns the server address for listening
func (a *App) GetAddress() string {
	return fmt.Sprintf("%s:%s", a.config.ServerHost, a.config.ServerPort)
}
