package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}

	return config
}
