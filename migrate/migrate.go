package main

import (
	"fmt"

	"github.com/leroytan/go-backend/initializers"
	"github.com/leroytan/go-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Role{})
	initializers.DB.AutoMigrate(&models.Course{})
	initializers.DB.AutoMigrate(&models.Category{})
	initializers.DB.AutoMigrate(&models.Subcategory{})
	initializers.DB.AutoMigrate(&models.PollsOptions{})
	initializers.DB.AutoMigrate(&models.PollsOptionsVotes{})

	roles := []*models.Role{
		{Title: "user", ID: 1},
		{Title: "admin", ID: 2},
	}

	initializers.DB.Create(roles)
	admin := models.User{Email: "admin@gmail.com", Username: "admin", Password: string(hashpassword("123456")), RoleID: 2}

	initializers.DB.Create(&admin)

}

//Hash password

func hashpassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		fmt.Println("Failed to hash")
	}
	return hash
}
