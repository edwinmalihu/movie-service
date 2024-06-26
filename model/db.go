package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {
	USER := os.Getenv("POSTGRES_USER")
	PASS := os.Getenv("POSTGRES_PASSWORD")
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("POSTGRES_DB")
	SSLMODE := os.Getenv("SSLMODE")
	TIMEZONE := os.Getenv("TIMEZONE")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", HOST, USER, PASS, DBNAME, PORT, SSLMODE, TIMEZONE)

	return gorm.Open(postgres.Open(url), &gorm.Config{Logger: newLogger})

}
