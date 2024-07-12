package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/meilisearch"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/services"
)

type JobHandler interface {
	GetJobs(c *gin.Context)
	GetJob(c *gin.Context)
}

type JobHandlerImpl struct {
	service services.JobService
}

func NewJobHandler(service services.JobService) *JobHandlerImpl {
	return &JobHandlerImpl{service: service}
}

func (h *JobHandlerImpl) GetJobs(c *gin.Context) {
	// handle get jobs logic
	jobs, err := h.service.GetJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobs,
	})
}

func (h *JobHandlerImpl) GetJob(c *gin.Context) {
	// Convert param id to int64
	idStr := c.Param("id")

	idInt64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle the error appropriately, e.g., return a 400 Bad Request response
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// handle get job logic
	job, err := h.service.GetJob(idInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": job,
	})
}

type SearchQuery struct {
	Query string `json:"query"`
}

func (h *JobHandlerImpl) SearchJobs(c *gin.Context) {
	// Parse body request to SearchQuery model
	searchQuery := SearchQuery{}
	err := c.ShouldBindJSON(&searchQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Validate SearchQuery
	if searchQuery.Query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	// handle search job logic
	jobs, err := meilisearch.SearchDocumentsInIndex("Jobs", searchQuery.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobs,
	})
}
