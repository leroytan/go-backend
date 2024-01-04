package main

import (
	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Course{})
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.Subcategory{})
	initializers.DB.AutoMigrate(&models.PollsOptions{})
	initializers.DB.AutoMigrate(&models.PollsOptionsVotes{})

}
