package config

import (
	"fmt"
	"incident-ai-backend/logger"
	"incident-ai-backend/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error
	if os.Getenv("TESTING") == "true" {
		err = godotenv.Load("./.env")
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		logger.Logger.Error("Error loading .env file", err)
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/incident_db?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Error("DB connection failed: %v", err)
		return nil
	}

	db.AutoMigrate(&models.Incident{})
	return db
}
