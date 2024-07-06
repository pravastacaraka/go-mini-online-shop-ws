package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/utils"
)

func NewDatabase(cfg *viper.Viper) *gorm.DB {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbHost := cfg.GetString("database.postgres.host")
	if os.Getenv("POSTGRES_HOST") != "" {
		dbHost = os.Getenv("POSTGRES_HOST")
	}

	dbPort := cfg.GetInt("database.postgres.port")
	if os.Getenv("POSTGRES_PORT") != "" {
		dbPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	}

	sslMode := "disable"
	if utils.IsProduction() {
		sslMode = "require"
	}

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s TimeZone=UTC", dbHost, dbPort, dbName, dbUser, dbPass, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}

	connection.SetMaxIdleConns(viper.GetInt("database.postgres.pool.idle"))
	connection.SetMaxOpenConns(viper.GetInt("database.postgres.pool.max"))
	connection.SetConnMaxLifetime(time.Second * time.Duration(viper.GetInt("database.postgres.pool.lifetime")))

	return db
}
