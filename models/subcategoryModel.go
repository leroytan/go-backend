package models

import "gorm.io/gorm"

type Subcategory struct {
	gorm.Model
	Title       string
	Description string
	Posts       []Post
	CategoryID  uint
}
