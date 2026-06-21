package main

import (
	"ecom/internal/shared/config"
	"ecom/internal/shared/logger"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	logX := logger.New(cfg.LogLevel, cfg.LogFormat)
	logX.Info("INITIALIZED")
}
