package repositories

import (
	"context"

	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobRepository interface {
	GetJobs() ([]models.Job, error)
	GetJob(id int64) (models.Job, error)
}

type JobRepositoryImpl struct {
	db *mongo.Client
}

func NewJobRepository(db *mongo.Client) *JobRepositoryImpl {
	return &JobRepositoryImpl{db: db}
}

func (r *JobRepositoryImpl) GetJobs() ([]models.Job, error) {
	var jobs []models.Job

	cursor, err := r.db.Database(configs.EnvDatabaseName()).Collection("Jobs").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *JobRepositoryImpl) GetJob(id int64) (models.Job, error) {
	filter := bson.M{"_id": id}

	var job models.Job

	jobDoc := r.db.Database(configs.EnvDatabaseName()).Collection("Jobs").FindOne(context.Background(), filter)

	err := jobDoc.Decode(&job)
	if err != nil {
		return models.Job{}, err
	}

	return job, err

}

func (r *JobRepositoryImpl) SearchJobs(query string) ([]models.Job, error) {
	var jobs []models.Job

	filter := bson.M{"$text": bson.M{"$search": query}}

	cursor, err := r.db.Database(configs.EnvDatabaseName()).Collection("Jobs").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}
