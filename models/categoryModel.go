package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title         string
	Description   string
	Subcategories []Subcategory
}
