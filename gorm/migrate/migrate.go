package main

import (
	"github.com/feynmaz/go/gorm/initializers"
	"github.com/feynmaz/go/gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
