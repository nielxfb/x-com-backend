package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Type        string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	OtherUserID uint
	OtherUser   User `gorm:"foreignKey:OtherUserID"`
	ThreadID    uint
	Thread      Thread `gorm:"foreignKey:ThreadID"`
	IsRead      bool
}
