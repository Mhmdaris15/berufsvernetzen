package repositories

import (
	"context"
	"log"

	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SurveyRepository interface {
	GetSurveys() ([]models.Survey, error)
	GetSurvey(id string) (models.Survey, error)
	CreateSurvey(survey models.Survey) (models.Survey, error)
	UpdateSurvey(survey models.Survey) (models.Survey, error)
	DeleteSurvey(id string) (models.Survey, error)
	PostJsonSurvey(survey map[string]interface{}) (map[string]interface{}, error)
	PostJsonSurveys(surveys []map[string]interface{}) ([]map[string]interface{}, error)
	GetFeedbacks() ([]string, error)
}

type SurveyRepositoryImpl struct {
	db *mongo.Client
}

func NewSurveyRepository(db *mongo.Client) *SurveyRepositoryImpl {
	return &SurveyRepositoryImpl{db: db}
}

func (r *SurveyRepositoryImpl) GetSurveys() ([]models.Survey, error) {
	var surveys []models.Survey

	cursor, err := r.db.Database(configs.EnvDatabaseName()).Collection("surveys").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &surveys); err != nil {
		return nil, err
	}

	return surveys, nil
}

func (r *SurveyRepositoryImpl) GetSurvey(id string) (models.Survey, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Survey{}, err
	}
	filter := bson.M{"_id": objectID}

	var survey models.Survey

	surveyDoc := r.db.Database(configs.EnvDatabaseName()).Collection("surveys").FindOne(context.Background(), filter)

	err = surveyDoc.Decode(&survey)
	if err != nil {
		return models.Survey{}, err
	}

	return survey, err
}

func (r *SurveyRepositoryImpl) CreateSurvey(survey models.Survey) (models.Survey, error) {
	insertedSurvey, err := r.db.Database(configs.EnvDatabaseName()).Collection("surveys").InsertOne(context.Background(), survey)
	if err != nil {
		return models.Survey{}, err
	}

	filter := bson.M{"_id": insertedSurvey.InsertedID}

	var createdSurvey models.Survey

	surveyDoc := r.db.Database(configs.EnvDatabaseName()).Collection("surveys").FindOne(context.Background(), filter)

	err = surveyDoc.Decode(&createdSurvey)
	if err != nil {
		return models.Survey{}, err
	}

	return createdSurvey, err
}

func (r *SurveyRepositoryImpl) UpdateSurvey(survey models.Survey) (models.Survey, error) {
	filter := bson.M{"_id": survey.ID}

	update := bson.M{
		"$set": bson.M{
			"current_activity": survey.CurrentActivity,
			"details":          survey.Details,
			"satisfaction":     survey.Satisfaction,
			"suggestions":      survey.Suggestions,
		},
	}

	_, err := r.db.Database(configs.EnvDatabaseName()).Collection("surveys").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.Survey{}, err
	}

	return survey, nil
}

func (r *SurveyRepositoryImpl) DeleteSurvey(id string) (models.Survey, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Survey{}, err
	}

	filter := bson.M{"_id": objectID}

	var survey models.Survey

	surveyDoc := r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").FindOneAndDelete(context.Background(), filter)

	err = surveyDoc.Decode(&survey)
	if err != nil {
		return models.Survey{}, err
	}

	return survey, err
}
func (r *SurveyRepositoryImpl) PostJsonSurvey(survey map[string]interface{}) (map[string]interface{}, error) {
	insertedSurvey, err := r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").InsertOne(context.Background(), survey)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertedSurvey.InsertedID}

	var createdSurvey map[string]interface{}

	err = r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").FindOne(context.Background(), filter).Decode(&createdSurvey)
	if err != nil {
		return nil, err
	}

	return createdSurvey, nil
}

func (r *SurveyRepositoryImpl) PostJsonSurveys(surveys []map[string]interface{}) ([]map[string]interface{}, error) {
	var createdSurveys []map[string]interface{}

	for _, survey := range surveys {
		insertedSurvey, err := r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").InsertOne(context.Background(), survey)
		if err != nil {
			return nil, err
		}

		filter := bson.M{"_id": insertedSurvey.InsertedID}

		var createdSurvey map[string]interface{}

		err = r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").FindOne(context.Background(), filter).Decode(&createdSurvey)
		if err != nil {
			return nil, err
		}

		createdSurveys = append(createdSurveys, createdSurvey)
	}

	return createdSurveys, nil
}

func (r *SurveyRepositoryImpl) GetFeedbacks() ([]string, error) {
	var surveys []map[string]interface{}

	cursor, err := r.db.Database(configs.EnvDatabaseName()).Collection("Surveys").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &surveys); err != nil {
		return nil, err
	}

	var feedbacks []string

	for _, survey := range surveys {
		if suggestion, ok := survey["improvement-suggestions"]; ok {
			if str, ok := suggestion.(string); ok && str != "" {
				feedbacks = append(feedbacks, str)
			}
		}
	}

	// Extract unique feedbacks and remove duplications
	uniqueFeedbacks := RemoveDuplicates(feedbacks)

	log.Printf("Unique Feedbacks: %v", uniqueFeedbacks)

	return uniqueFeedbacks, nil
}

// Helper function to remove duplicates from a slice of strings
func RemoveDuplicates(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, item := range strSlice {
		if _, value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}
	return list
}
