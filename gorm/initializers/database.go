package initializers

import (
	"log"
	"os"

	"github.com/feynmaz/go/gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
}

func CreateRelations() {
	DB.AutoMigrate(&models.Post{})
}
