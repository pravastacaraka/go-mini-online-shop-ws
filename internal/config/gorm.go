package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *viper.Viper) *gorm.DB {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := cfg.GetInt32("database.postgres.port")

	dsn := fmt.Sprintf("host=postgres port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=UTC", dbPort, dbName, dbUser, dbPass)

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
