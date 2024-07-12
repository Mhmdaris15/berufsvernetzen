package seeder

import (
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
)

func Seed() ([]models.User, []models.Survey, error) {
	// Seed User
	users, err := SeedUser()
	if err != nil {
		return nil, nil, err
	}

	// Seed Survey
	surveys, err := SeedSurvey()
	if err != nil {
		return users, nil, err
	}

	return users, surveys, nil
}
