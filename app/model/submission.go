package model

import "github.com/jinzhu/gorm"

type Submission struct {
	gorm.Model
	UserId      uint
	ChallengeId uint
	Points      int
	Timestamp   string
}
