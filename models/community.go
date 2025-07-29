package models

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	CommunityName string
	Description   string
	MemberCount   int
	LogoPath      *string
	BannerPath    *string
	CategoryID    uint
	Category      CommunityCategory `gorm:"foreignKey:CategoryID"`
	Rules         []CommunityRule   `gorm:"foreignKey:CommunityID"`
	Members       []CommunityMember `gorm:"foreignKey:CommunityID"`
	Threads       []*Thread         `gorm:"foreignKey:CommunityID"`
}

type CommunityCategory struct {
	gorm.Model
	CategoryName string `gorm:"uniqueIndex"`
}

type CommunityMember struct {
	gorm.Model
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	CommunityID uint
	Community   Community `gorm:"foreignKey:CommunityID"`
	IsModerator bool
}

type CommunityRule struct {
	gorm.Model
	CommunityID uint
	Community   Community `gorm:"foreignKey:CommunityID"`
	Number      int
	RuleHeader  string
	RuleContent string
}
