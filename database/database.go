package database

import (
	"github.com/rombintu/gopasswd.git/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Users{}, &models.Passwords{})
	return db
}

func Get_db() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
