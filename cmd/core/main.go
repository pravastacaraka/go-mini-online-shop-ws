package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/config"
)

func main() {
	cfg := config.NewViper()
	db := config.NewDatabase(cfg)
	validator := config.NewValidator()
	app := config.NewFiber(cfg)

	config.Bootstrap(&config.BootstrapConfig{
		Config:    cfg,
		DB:        db,
		Validator: validator,
		App:       app,
	})

	port := 8080
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server on port :%d", port)
	}
}
