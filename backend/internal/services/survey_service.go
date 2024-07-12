package services

import (
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
)

type SurveyService interface {
	GetSurveys() ([]models.Survey, error)
	GetSurvey(id string) (models.Survey, error)
	CreateSurvey(survey models.Survey) (models.Survey, error)
	UpdateSurvey(survey models.Survey) (models.Survey, error)
	DeleteSurvey(id string) (models.Survey, error)
	PostJsonSurvey(survey map[string]interface{}) (map[string]interface{}, error)
	PostJsonSurveys(surveys []map[string]interface{}) ([]map[string]interface{}, error)
	GetFeedbacks() ([]string, error)
}

type SurveyServiceImpl struct {
	repo repositories.SurveyRepository
}

func NewSurveyService(repo repositories.SurveyRepository) *SurveyServiceImpl {
	return &SurveyServiceImpl{repo: repo}
}

func (s *SurveyServiceImpl) GetSurveys() ([]models.Survey, error) {
	return s.repo.GetSurveys()
}

func (s *SurveyServiceImpl) GetSurvey(id string) (models.Survey, error) {
	return s.repo.GetSurvey(id)
}

func (s *SurveyServiceImpl) CreateSurvey(survey models.Survey) (models.Survey, error) {
	return s.repo.CreateSurvey(survey)
}

func (s *SurveyServiceImpl) UpdateSurvey(survey models.Survey) (models.Survey, error) {
	return s.repo.UpdateSurvey(survey)
}

func (s *SurveyServiceImpl) DeleteSurvey(id string) (models.Survey, error) {
	return s.repo.DeleteSurvey(id)
}

func (s *SurveyServiceImpl) PostJsonSurvey(survey map[string]interface{}) (map[string]interface{}, error) {
	return s.repo.PostJsonSurvey(survey)
}

func (s *SurveyServiceImpl) PostJsonSurveys(surveys []map[string]interface{}) ([]map[string]interface{}, error) {
	return s.repo.PostJsonSurveys(surveys)
}

func (s *SurveyServiceImpl) GetFeedbacks() ([]string, error) {
	return s.repo.GetFeedbacks()
}
