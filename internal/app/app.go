package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevich/blockchain/internal/config"
	"github.com/vashkevich/blockchain/pkg/logger"
)

// App represents the main application
type App struct {
	config *config.Config
	logger *logger.Logger
	db     *pgxpool.Pool
}

// New creates a new application
func New(cfg *config.Config, log *logger.Logger, db *pgxpool.Pool) (*App, error) {
	app := &App{
		config: cfg,
		logger: log,
		db:     db,
	}

	return app, nil
}

// Run starts the application
func (a *App) Run(ctx context.Context) error {
	a.logger.Info("Application started")

	// TODO: Start HTTP server and workers

	<-ctx.Done()
	return ctx.Err()
}

// Shutdown gracefully stops the application
func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("Shutting down application...")

	if a.db != nil {
		a.db.Close()
		a.logger.Info("Database connection pool closed")
	}

	// TODO: Shutdown HTTP server and workers

	return nil
}

// GetAddress returns the server address for listening
func (a *App) GetAddress() string {
	return fmt.Sprintf("%s:%s", a.config.ServerHost, a.config.ServerPort)
}
