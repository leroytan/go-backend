package main

import (
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
