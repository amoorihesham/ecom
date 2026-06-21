package main

import (
	"context"
	"ecom/internal/shared/config"
	"ecom/internal/shared/database"
	"ecom/internal/shared/logger"
	"fmt"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	logX.Info("", "db", db.Stats())

}
