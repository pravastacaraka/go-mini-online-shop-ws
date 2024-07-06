package config

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
)

func NewRedis(cfg *viper.Viper) *redis.Storage {
	var (
		url, host, username, password string
		port                          int
	)

	if os.Getenv("REDIS_URL") != "" {
		url = os.Getenv("REDIS_URL")
	} else {
		host = cfg.GetString("database.redis.host")
		if os.Getenv("REDIS_HOST") != "" {
			host = os.Getenv("REDIS_HOST")
		}

		port = cfg.GetInt("database.redis.port")
		if os.Getenv("REDIS_PORT") != "" {
			port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
		}

		if os.Getenv("REDIS_USERNAME") != "" {
			username = os.Getenv("REDIS_USERNAME")
		}

		if os.Getenv("REDIS_PASSWORD") != "" {
			password = os.Getenv("REDIS_PASSWORD")
		}
	}

	store := redis.New(redis.Config{
		URL:      url,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		PoolSize: cfg.GetInt("database.redis.pool") * runtime.GOMAXPROCS(0),
	})

	return store
}
