package models

import "gorm.io/gorm"

type Hashtag struct {
	gorm.Model
	HashtagName string `gorm:"uniqueIndex"`
	Count       int
}
