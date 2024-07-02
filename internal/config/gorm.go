package config

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *viper.Viper) *gorm.DB {
	dbHost := cfg.GetString("database.postgres.host")
	dbPort := cfg.GetInt32("database.postgres.port")
	dbUser := cfg.GetString("database.postgres.username")
	dbPass := cfg.GetString("database.postgres.password")
	dbName := cfg.GetString("database.postgres.name")

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=UTC", dbHost, dbPort, dbName, dbUser, dbPass)

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
