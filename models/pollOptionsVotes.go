package models

import "gorm.io/gorm"

type PollsOptionsVotes struct {
	gorm.Model
	PollsOptionsID uint
	UserID         uint
}
