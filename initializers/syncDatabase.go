package initializers

import (
	"github.com/leroytan/go-backend/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
