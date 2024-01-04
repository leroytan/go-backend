package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string `binding:"required"`
	Content       string `binding:"required"`
	ParentpostID  *uint
	Parentpost    *Post
	UserID        uint `binding:"required"`
	SubcategoryID uint `binding:"required"`
	PollsOptions  []PollsOptions
}
