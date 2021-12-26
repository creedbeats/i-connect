package database

import (
	"fmt"
	"log"
	"os"

	"github.com/creedbeats/i-connect.git/api/config"
	"github.com/creedbeats/i-connect.git/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error
	DbHost, DbPort, DbUser, DbName, DbPassword := config.Get("DB_HOST"), config.Get("DB_PORT"), config.Get("DB_USER"), config.Get("DB_NAME"), config.Get("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// Add Migrations
	DB.AutoMigrate(&models.User{})
	log.Println("Database Migrated")
}