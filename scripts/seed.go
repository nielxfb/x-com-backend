package scripts

import (
	"github.com/nielxfb/x-com-backend/database"
	"github.com/nielxfb/x-com-backend/models"
)

func SeedDatabase() {
	db := database.GetConnection()

	securityQuestions := []string {
		"What was the name of your first pet?",
		"What city were you born in?",
		"WHat is your favorite video game?",
		"What was the name of your first school?",
		"What was your childhood nickname?",
	}

	for _, question := range securityQuestions {
		securityQuestion := models.SecurityQuestion{
			Question: question,
		}

		db.Create(&securityQuestion)
	}
}