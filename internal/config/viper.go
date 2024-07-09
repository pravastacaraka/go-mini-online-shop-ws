package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/utils"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigType("json")

	if utils.IsProduction() {
		config.SetConfigName("config.prod")
	} else {
		config.SetConfigName("config.local")
		config.AddConfigPath("./../../") // Local development path
	}
	config.AddConfigPath(".") // Add current directory as a fallback

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}

	return config
}
