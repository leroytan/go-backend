package models

import "gorm.io/gorm"

type PollsOptions struct {
	gorm.Model
	Title             string
	PostID            uint
	PollsOptionsVotes []PollsOptionsVotes
}
