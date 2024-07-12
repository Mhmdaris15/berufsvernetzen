package meilisearch

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
)

var (
	Client *meilisearch.Client
)

func ConnectToMeiliSearch() *meilisearch.Client {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: configs.EnvMeiliMasterKey(),
	})

	Client = client

	return client
}

func InitializeIndex(name string) meilisearch.IndexInterface {
	index := Client.Index(name)

	return index
}

func AddDocumentsToIndex(index meilisearch.IndexInterface) (interface{}, error) {

	// Initialize Job Repository
	jobRepository := repositories.NewJobRepository(mongodb.DB)
	jobs, err := jobRepository.GetJobs()
	if err != nil {
		return nil, err
	}

	res, err := index.AddDocuments(jobs)
	if err != nil {
		return res, err
	}

	// Print Response
	log.Printf("AddDocumentsToIndex Response: %v", res)

	return res, nil
}

func DeleteDocumentsFromIndex(index meilisearch.IndexInterface) (interface{}, error) {
	res, err := index.DeleteAllDocuments()
	if err != nil {
		return res, err
	}

	// Print Response
	log.Printf("DeleteDocumentsFromIndex Response: %v", res)

	return res, nil
}

func SearchDocumentsInIndex(indexName string, query string) (interface{}, error) {
	index, err := Client.GetIndex(indexName)
	if err != nil {
		log.Printf("Error while getting index: %v", err)
		return nil, err
	}

	searchRequest := &meilisearch.SearchRequest{
		Query: query,
		Limit: 10,
	}

	res, err := index.Search(query, searchRequest)
	if err != nil {
		log.Printf("Error while searching index: %v", err)
		return res, err
	}

	return res, nil
}
