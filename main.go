package main

import (
	"github.com/nielxfb/x-com-backend/database"
	"github.com/nielxfb/x-com-backend/models"
	"github.com/nielxfb/x-com-backend/scripts"
)

func main() {
	db := database.GetConnection()
	db.AutoMigrate(&models.SecurityQuestion{})
	scripts.SeedDatabase()
}
