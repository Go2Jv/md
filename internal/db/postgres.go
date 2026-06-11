package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	return db, err
}
