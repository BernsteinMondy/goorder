package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/BernsteinMondy/goorder/orders/src/internal/domain"
	"github.com/BernsteinMondy/goorder/orders/src/internal/impl"
	"github.com/BernsteinMondy/goorder/orders/src/pkg/database"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		println("run() returned error: " + err.Error())
	}
}

func run() (err error) {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)
	defer cancel()

	slog.Info("Loading config")
	cfg, err := loadConfigFromEnv()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	slog.Info("Config loaded")

	select {
	case <-ctx.Done():
	default:
	}

	slog.Info("Connecting to database")
	db, err := NewDatabase(cfg)
	if err != nil {
		return fmt.Errorf("new SQL: %w", err)
	}
	slog.Info("Connected to database")
	defer func() {
		slog.Info("Closing database connection")
		err = db.Close()
		if err != nil {
			err = errors.Join(err, fmt.Errorf("close database connection: %w", err))
		}
		slog.Info("Database connection closed")
	}()

	ordersRepo := impl.NewOrderRepository(db)
	productsRepo := impl.NewProductRepository(db)

	service := domain.NewService(
		ordersRepo,
		productsRepo,
	)

	_ = service

	<-ctx.Done()
	return nil
}

func NewDatabase(c *Config) (*sql.DB, error) {
	dbCfg := database.Config{}

	db, err := database.NewSQL(dbCfg)
	if err != nil {
		return nil, fmt.Errorf("NewSQL() error: %w", err)
	}

	return db, nil
}
