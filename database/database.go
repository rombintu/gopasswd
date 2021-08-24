package database

import (
	"os"

	"github.com/rombintu/gopasswd.git/models"
	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// export CREDS="host=127.0.0.1 user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
	creds := os.Getenv("CREDS")
	db, err := gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Users{}, &models.Passwords{})
	return db
}

func Get_db() *gorm.DB {
	creds := os.Getenv("CREDS")
	db, err := gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
