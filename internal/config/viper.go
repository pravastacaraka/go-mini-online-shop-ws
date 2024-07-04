package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	env := os.Getenv("APP_ENV")

	config := viper.New()
	config.SetConfigType("json")

	if env == "production" {
		config.SetConfigName("config.prod")
		config.AddConfigPath("/usr/src/app/") // Docker container path
	} else {
		config.SetConfigName("config.local")
		config.AddConfigPath("./../") // Local development path
		config.AddConfigPath(".")     // Add current directory as a fallback
	}

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}

	return config
}
