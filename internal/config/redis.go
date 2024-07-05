package config

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
)

func NewRedis(cfg *viper.Viper) *redis.Storage {
	url := ""
	if os.Getenv("REDIS_URL") != "" {
		url = os.Getenv("REDIS_URL")
	}

	host := cfg.GetString("database.redis.host")
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}

	port := cfg.GetInt("database.redis.port")
	if os.Getenv("REDIS_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	}

	store := redis.New(redis.Config{
		Host:     host,
		URL:      url,
		Port:     port,
		PoolSize: cfg.GetInt("database.redis.pool") * runtime.GOMAXPROCS(0),
	})

	return store
}
