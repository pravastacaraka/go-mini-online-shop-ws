package main

import (
	"os"

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

	port := os.Getenv("PORT")
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server on port: %s, err: %s", port, err.Error())
	}
}
