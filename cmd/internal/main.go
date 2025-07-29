package main

import (
    "github.com/nielxfb/x-com-backend/database"
    "github.com/nielxfb/x-com-backend/models"
    "github.com/nielxfb/x-com-backend/scripts"
)

func main() {
    db := database.GetConnection()

    // Drop all tables in reverse dependency order
    db.Migrator().DropTable(&models.PollVote{})
    db.Migrator().DropTable(&models.PollOption{})
    db.Migrator().DropTable(&models.ThreadFile{})
    db.Migrator().DropTable(&models.ThreadLike{})
    db.Migrator().DropTable(&models.Thread{})
    db.Migrator().DropTable(&models.ThreadCategory{})
    db.Migrator().DropTable(&models.GroupMessage{})
    db.Migrator().DropTable(&models.GroupMember{})
    db.Migrator().DropTable(&models.Group{})
    db.Migrator().DropTable(&models.PrivateMessage{})
    db.Migrator().DropTable(&models.Notification{})
    db.Migrator().DropTable(&models.Report{})
    db.Migrator().DropTable(&models.CommunityRule{})
    db.Migrator().DropTable(&models.CommunityMember{})
    db.Migrator().DropTable(&models.Community{})
    db.Migrator().DropTable(&models.CommunityCategory{})
    db.Migrator().DropTable(&models.UserSecurityAnswer{})
    db.Migrator().DropTable(&models.FollowList{})
    db.Migrator().DropTable(&models.BlockList{})
    db.Migrator().DropTable(&models.User{})
    db.Migrator().DropTable(&models.SecurityQuestion{})
    db.Migrator().DropTable(&models.Hashtag{})

    // Create tables in dependency order
    // Base models without foreign key dependencies
    db.AutoMigrate(&models.SecurityQuestion{})
    db.AutoMigrate(&models.Hashtag{})
    db.AutoMigrate(&models.CommunityCategory{})
    db.AutoMigrate(&models.ThreadCategory{})
    db.AutoMigrate(&models.User{})

    // Models that depend on User
    db.AutoMigrate(&models.UserSecurityAnswer{})
    db.AutoMigrate(&models.BlockList{})
    db.AutoMigrate(&models.FollowList{})
    db.AutoMigrate(&models.Report{})
    db.AutoMigrate(&models.PrivateMessage{})
    db.AutoMigrate(&models.Group{})
    db.AutoMigrate(&models.GroupMember{})
    db.AutoMigrate(&models.GroupMessage{})

    // Models that depend on CommunityCategory
    db.AutoMigrate(&models.Community{})
    db.AutoMigrate(&models.CommunityMember{})
    db.AutoMigrate(&models.CommunityRule{})

    // Models that depend on multiple entities
    db.AutoMigrate(&models.Thread{})
    db.AutoMigrate(&models.ThreadLike{})
    db.AutoMigrate(&models.ThreadFile{})
    db.AutoMigrate(&models.Notification{})
    db.AutoMigrate(&models.PollOption{})
    db.AutoMigrate(&models.PollVote{})

    scripts.SeedDatabase()
}