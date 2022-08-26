package db

import (
	"github.com/YelloJW/simple-rest-api/app/packages/db/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Connect() {
  database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&models.User{})

  DB = database
}

