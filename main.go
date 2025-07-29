package main

import (
	"github.com/nielxfb/x-com-backend/database"
	"github.com/nielxfb/x-com-backend/models"
)

func main() {
	db := database.GetConnection()
	db.AutoMigrate(&models.User{})
}
