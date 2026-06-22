package main

import (
	"context"
	"ecom/cmd/application"
	"ecom/internal/modules/auth"
	"ecom/internal/modules/catalog"
	"ecom/internal/shared/config"
	"ecom/internal/shared/database"
	"ecom/internal/shared/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

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
		logX.Error("Database Error", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logX.Info("Apply db migrations...")
	if err := database.Migrate(ctx, cfg.DatabaseUrl); err != nil {
		logX.Error("Migration", "error", err.Error())
	}

	logX.Info("Initializing the mux routers...")
	mux := http.NewServeMux()
	catalog.Initialize(mux, db, logX)
	auth.Initialize(mux, db, logX)
	logX.Info("Initialized the mux routers...")

	app := application.NewApplication(&application.AppConfig{
		Addr: cfg.Addr,
		Mux:  mux,
	})

	logX.Info("server started", "host", cfg.Addr)
	app.Run(ctx, logX, errCh)

	select {
	case sig := <-sigCh:
		logX.Info("shutdown signal received.", "signal", sig.String())
		logX.Info("shuting down....")
	case err := <-errCh:
		logX.Error("server start faild", "error", err)
	}

	if err := app.Shutdown(ctx); err != nil {
		logX.Error("server shutdown faild", "error", err)
	}

	logX.Info("server shutdown complete")

}
