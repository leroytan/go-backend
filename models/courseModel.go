package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string
	Description string
	Categories  []Category
	Users       []User `gorm:"many2many:course_users;"`
}
