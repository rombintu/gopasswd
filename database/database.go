package database

import (
	"github.com/rombintu/gopasswd.git/models"
	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Users{}, &models.Passwords{})
	return db
}

func Get_db() *gorm.DB {
	dsn := "host=localhost user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
