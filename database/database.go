package database

import (
	"log"

	"github.com/glebarez/sqlite" // Pure Go SQLite driver that doesn't require cgo
	"github.com/yourusername/go-fiber-books/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }
    database.AutoMigrate(&models.Book{})
    DB = database
}
