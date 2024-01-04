package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title         string
	Description   string
	CourseID      uint
	Subcategories []Subcategory
}
