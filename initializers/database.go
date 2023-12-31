package initializers

import (
	"github.com/glebarez/sqlite" //Pure go SQLite driver not based on CGO
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	DB, err = gorm.Open(sqlite.Open("db/post.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
