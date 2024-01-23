package initializers

import (
	"github.com/leroytan/go-backend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.Course{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Subcategory{})
	DB.AutoMigrate(&models.PollsOptions{})
	DB.AutoMigrate(&models.PollsOptionsVotes{})

}
