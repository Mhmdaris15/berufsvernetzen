package seeder

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedSurvey() ([]models.Survey, error) {
	// Get all users
	users := []models.User{}
	userCollection := mongodb.GetCollection(mongodb.DB, "Users")
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		// log.Printf("Error when getting users 1: %v", err)
		return nil, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		// log.Printf("Error when getting users 2: %v", err)
		return nil, err
	}

	currentActivities := []string{"working", "enterpreneur", "study", "unemployed"}

	numOfSurveys := 10

	// Seed Survey
	surveys := []models.Survey{}
	surveyCollection := mongodb.GetCollection(mongodb.DB, "Surveys")
	for i := 0; i < numOfSurveys; i++ {
		survey := models.Survey{
			ID:              primitive.NewObjectID(),
			UserId:          users[i].ID.Hex(),
			CurrentActivity: currentActivities[i%4],
			Satisfaction:    gofakeit.RandomString([]string{"very satisfied", "satisfied", "neutral", "unsatisfied", "very unsatisfied"}),
			Suggestions:     gofakeit.Sentence(10),
		}

		switch survey.CurrentActivity {
		case "working":
			survey.Details = models.Working{
				CompanyName:            gofakeit.Company(),
				IndustrySector:         gofakeit.Hobby(),
				Location:               gofakeit.Address().Address,
				Position:               gofakeit.JobTitle(),
				Status:                 gofakeit.RandomString([]string{"permanent", "contract", "internship"}),
				GrossSalary:            gofakeit.Number(20000, 100000),
				WorkStudyConnectedness: gofakeit.RandomString([]string{"related", "unrelated"}),
				FirstJob:               gofakeit.Bool(),
				FirstDayWork:           gofakeit.Date().String(),
				IsAverageMinimumWage:   gofakeit.Bool(),
			}
		case "enterpreneur":
			survey.Details = models.Enterpreneurship{
				Product:          gofakeit.AppName(),
				NetIncomeAverage: gofakeit.Number(20000, 100000),
				Unit:             gofakeit.RandomString([]string{"month", "year"}),
				StartDate:        gofakeit.Date().String(),
			}

		case "study":
			survey.Details = models.FurtherStudy{
				StudyProgram: gofakeit.RandomString([]string{"S1", "S2", "S3"}),
				University:   gofakeit.Company(),
				Level:        gofakeit.RandomString([]string{"undergraduate", "graduate"}),
				StartDate:    gofakeit.Date().String(),
			}

		case "unemployed":
			survey.Details = models.NotWorking{
				AlreadyWorkingOrEnterpreneurship: gofakeit.Bool(),
				Reason:                           gofakeit.Sentence(10),
				StartWorkingOrEnterpreneurship:   gofakeit.Date().String(),
				EndWorkingOrEnterpreneurship:     gofakeit.Date().String(),
			}
		}

		surveys = append(surveys, survey)
	}

	// Insert surveys to surveyCollection
	surveyInterfaces := make([]interface{}, len(surveys))
	for i, survey := range surveys {
		surveyInterfaces[i] = survey
	}

	results, err := surveyCollection.InsertMany(context.TODO(), surveyInterfaces)
	if err != nil {
		log.Printf("Error when inserting surveys: %v", err)
		return nil, err
	}

	// If the number of inserted surveys is not equal to the number of surveys, return error
	if len(results.InsertedIDs) != len(surveys) {
		log.Printf("Error when inserting surveys: %v", err)
		return nil, err
	}

	return surveys, nil
}
