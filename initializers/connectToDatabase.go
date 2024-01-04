package initializers

import (
	"os"

	"github.com/glebarez/sqlite" //Pure go SQLite driver not based on CGO
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Connect to SQlite database
func ConnectToDB() {
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}
