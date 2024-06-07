package internal

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

func initializeDB() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_PORT == "" || DB_NAME == "" {
		err = fmt.Errorf("missing required environment variables for database connection")
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/London", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		err = fmt.Errorf("failed to connect to the database: %w", err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)
}

func CreateConnection() (*gorm.DB, error) {
	once.Do(initializeDB)

	if err != nil {
		return nil, err
	}

	return db, nil
}
