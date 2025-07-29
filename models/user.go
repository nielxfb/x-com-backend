package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	FollowerCount int
	FollowingCount int
	FullName string
	Email string
	Password string
	Gender string
	ProfilePicture string
	Banner string
	DateOfBirth time.Time
	SubscribeNewsletter bool
	VerificationCode string
	IsBanned bool
	IsActive bool
	IsVerified bool
	IsPrivate bool
}