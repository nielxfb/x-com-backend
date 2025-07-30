package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username            string
	FollowerCount       int
	FollowingCount      int
	FullName            string
	Email               string
	Password            string
	Gender              string
	ProfilePicture      *string
	Banner              *string
	DateOfBirth         time.Time
	SubscribeNewsletter bool
	VerificationCode    *string
	IsBanned            bool
	IsActive            bool
	IsVerified          bool
	IsPrivate           bool
	ThreadsLiked        []*ThreadLike        `gorm:"foreignKey:UserID"`
	SecurityAnswers     []UserSecurityAnswer `gorm:"foreignKey:UserID"`
	Notifications       []*Notification      `gorm:"foreignKey:UserID"`
	BlockedUsers        []*BlockList         `gorm:"foreignKey:UserID"`
	Following           []*FollowList        `gorm:"foreignKey:FollowerID"`
	Followers           []*FollowList        `gorm:"foreignKey:FollowingID"`
}

type SecurityQuestion struct {
	gorm.Model
	Question string `gorm:"uniqueIndex"`
}

type UserSecurityAnswer struct {
	gorm.Model
	UserID             uint
	User               User `gorm:"foreignKey:UserID"`
	SecurityQuestionID uint
	SecurityQuestion   SecurityQuestion `gorm:"foreignKey:SecurityQuestionID"`
	Answer             string
}

type BlockList struct {
	gorm.Model
	UserID        uint
	User          User `gorm:"foreignKey:UserID"`
	BlockedUserID uint
	BlockedUser   User `gorm:"foreignKey:BlockedUserID"`
}

type FollowList struct {
	gorm.Model
	FollowerID  uint
	Follower    User `gorm:"foreignKey:FollowerID"`
	FollowingID uint
	Following   User `gorm:"foreignKey:FollowingID"`
}
