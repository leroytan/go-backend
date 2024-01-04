package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string `binding:"required"`
	Content       string `binding:"required"`
	UserID        uint   `binding:"required"`
	SubcategoryID uint   `binding:"required"`
}
