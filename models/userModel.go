package models

import "gorm.io/gorm"

//User has many Posts, UserID is the foreign key
type User struct {
	gorm.Model
	Email    string `binding:"required" gorm:"unique;not null"`
	Username string `binding:"required" gorm:"not null"`
	Password string `binding:"required"`
	Posts    []Post
}
