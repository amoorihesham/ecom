package main

import (
	"context"
	"ecom/cmd/application"
	"ecom/internal/modules/auth"
	"ecom/internal/shared/config"
	"ecom/internal/shared/database"
	"ecom/internal/shared/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	errCh := make(chan error, 1)

	cfg, err := config.New(".env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logX := logger.New(cfg.LogLevel, cfg.LogFormat)

	logX.Info("Initializing Database...")
	db, err := database.NewDatabase(ctx, &database.DBConfig{
		ConnectionString: cfg.DatabaseUrl,
		MaxConns:         cfg.DBMaxConns,
		MinConns:         cfg.DBMinConns,
		PingTimeout:      cfg.PingTimeout,
	})
	if err != nil {
		logX.Error("Database Init", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logX.Info("Running db migrations...")
	if err := database.Migrate(context.Background(), cfg.DatabaseUrl); err != nil {
		logX.Error("Migration", "error", err.Error())
		os.Exit(1)
	}

	logX.Info("Initializing routes...")

	handler := application.NewRouter(db, logX, &application.RouterConfig{
		Auth: auth.AuthConfig{
			BcryptCost: cfg.BcryptCost,
			JWTSecret:  cfg.JWTSecret,
		}})

	app := application.NewApplication(&application.AppConfig{
		Addr:    cfg.Addr,
		Handler: handler,
	})

	app.Run(ctx, logX, errCh)

	select {
	case <-ctx.Done():
		logX.Info("shutdown signal received.")

	case err := <-errCh:
		logX.Error("server error occured", "error", err)
	}

	logX.Info("shutting down server")

	if err := app.Shutdown(context.Background()); err != nil {
		logX.Error("shutdown failed", "error", err)
		os.Exit(1)
	}

	logX.Info("server shutdown complete")

}
