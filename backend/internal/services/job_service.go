package services

import (
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
)

type JobService interface {
	GetJobs() ([]models.Job, error)
	GetJob(id int64) (models.Job, error)
}

type JobServiceImpl struct {
	repo repositories.JobRepository
}

func NewJobService(repo repositories.JobRepository) *JobServiceImpl {
	return &JobServiceImpl{repo: repo}
}

func (s *JobServiceImpl) GetJobs() ([]models.Job, error) {
	return s.repo.GetJobs()
}

func (s *JobServiceImpl) GetJob(id int64) (models.Job, error) {
	return s.repo.GetJob(id)
}
