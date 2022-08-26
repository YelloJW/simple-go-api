package db

// TODO: shared package shouldn't dep on module
import (
	"github.com/YelloJW/simple-rest-api/app/users/infrastructure/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() {
  database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&models.User{})

  DB = database
}

