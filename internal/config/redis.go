package config

import (
	"os"
	"runtime"

	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
)

func NewRedis(cfg *viper.Viper) *redis.Storage {
	host := cfg.GetString("database.redis.host")

	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}

	store := redis.New(redis.Config{
		Host:     host,
		Port:     cfg.GetInt("database.redis.port"),
		PoolSize: cfg.GetInt("database.redis.pool") * runtime.GOMAXPROCS(0),
	})

	return store
}
