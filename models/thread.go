package models

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	LikeCount   int
	ReplyCount  int
	RepostCount int
	CommunityID *uint
	Community   *Community `gorm:"foreignKey:CommunityID"`
	CategoryID  *uint
	Category    *ThreadCategory `gorm:"foreignKey:CategoryID"`
	HashtagID   *uint
	Hashtag     *Hashtag `gorm:"foreignKey:HashtagID"`
	Content     string
	ReplyID     *uint
	Reply       *Thread `gorm:"foreignKey:ReplyID"`
	RepostID    *uint
	Repost      *Thread `gorm:"foreignKey:RepostID"`
	ScheduledAt *time.Time
	IsReply     bool
	IsRepost    bool
	IsPoll      bool
	IsAds       bool
	IsScheduled bool
	IsMedia     bool
	ThreadFiles []*ThreadFile `gorm:"foreignKey:ThreadID"`
	ThreadLikes []*ThreadLike `gorm:"foreignKey:ThreadID"`
	PollOptions []*PollOption `gorm:"foreignKey:ThreadID"`
	CreatorID   uint
	Creator     User `gorm:"foreignKey:CreatorID"`
}

type ThreadLike struct {
	gorm.Model
	ThreadID uint
	Thread   Thread `gorm:"foreignKey:ThreadID"`
	UserID   uint
	User     User `gorm:"foreignKey:UserID"`
}

type ThreadCategory struct {
	gorm.Model
	CategoryName string `gorm:"uniqueIndex"`
}

type ThreadFile struct {
	gorm.Model
	ThreadID uint
	Thread   Thread `gorm:"foreignKey:ThreadID"`
}

type PollOption struct {
	gorm.Model
	ThreadID  uint
	Thread    Thread `gorm:"foreignKey:ThreadID"`
	Option    string
	VoteCount int
}

type PollVote struct {
	gorm.Model
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	PollOptionID uint
	PollOption   PollOption `gorm:"foreignKey:PollOptionID"`
}
