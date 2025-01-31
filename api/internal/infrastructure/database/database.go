package database

import (
	"fmt"
	"os"
	"porsche-api/internal/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
		os.Getenv("HA_PROXY_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("HA_PROXY_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		panic(fmt.Sprintf("warning: unable to create extension uuid-ossp: %v", err))
	}

	if err := DB.AutoMigrate(&model.Car{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	SeedCar()
}
