package models

import "gorm.io/gorm"

type PrivateMessage struct {
	gorm.Model
	Message    string
	SenderID   uint
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint
	Receiver   User `gorm:"foreignKey:ReceiverID"`
	IsFile     bool
	FilePath   string
	IsRead     bool
}

type Group struct {
	gorm.Model
	Name         string
	GroupMembers []GroupMember   `gorm:"foreignKey:GroupID"`
	Messages     []*GroupMessage `gorm:"foreignKey:GroupID"`
}

type GroupMember struct {
	gorm.Model
	GroupID uint
	Group   Group `gorm:"foreignKey:GroupID"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

type GroupMessage struct {
	gorm.Model
	Message  string
	SenderID uint
	Sender   User `gorm:"foreignKey:SenderID"`
	GroupID  uint
	Group    Group `gorm:"foreignKey:GroupID"`
	IsFile   bool
	FilePath string
	IsRead   bool
}
