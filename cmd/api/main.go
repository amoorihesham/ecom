package main

import (
	"ecom/internal/config"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.New(".env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Print(cfg)
}
