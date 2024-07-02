package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/route"
)

type BootstrapConfig struct {
	Config *viper.Viper
	DB     *gorm.DB
	App    *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	route := route.RouteConfig{
		Config: config.Config,
		DB:     config.DB,
		App:    config.App,
	}
	route.Setup()
}
