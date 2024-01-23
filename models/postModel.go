package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title        string `binding:"required"`
	Content      string `binding:"required"`
	ParentpostID *uint
	Parentpost   *Post
	UserID       uint `binding:"required"`
	//Username string
	SubcategoryID uint `binding:"required"`
	PollsOptions  []PollsOptions
	//Userpollsvotes []PollsOptionsVotes
	Upvotecount   uint
	Downvotecount uint
	Comments      []*Post `gorm:"many2many:post_comments"`
}
