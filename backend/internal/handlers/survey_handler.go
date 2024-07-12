package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
)

type SurveyHandler interface {
	GetSurveys(c *gin.Context)
	GetSurvey(c *gin.Context)
	CreateSurvey(c *gin.Context)
	UpdateSurvey(c *gin.Context)
	DeleteSurvey(c *gin.Context)
	PostSurveyJson(c *gin.Context)
	PostSurveyJsons(c *gin.Context)
	GetFeedbacks(c *gin.Context)
}

type SurveyHandlerImpl struct {
	service repositories.SurveyRepository
}

func NewSurveyHandler(service repositories.SurveyRepository) *SurveyHandlerImpl {
	return &SurveyHandlerImpl{service: service}
}

func (h *SurveyHandlerImpl) GetSurveys(c *gin.Context) {
	surveys, err := h.service.GetSurveys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": surveys,
	})
}

func (h *SurveyHandlerImpl) GetSurvey(c *gin.Context) {
	survey, err := h.service.GetSurvey(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": survey,
	})
}

func (h *SurveyHandlerImpl) CreateSurvey(c *gin.Context) {
	survey := models.Survey{}

	// Check current_activity of survey either working, enterpreneur, study, or unemployed
	switch c.PostForm("current_activity") {
	case "bekerja":
		survey.Details = models.Working{}
	case "berwirausaha":
		survey.Details = models.Enterpreneurship{}
	case "kuliah":
		survey.Details = models.FurtherStudy{}
	case "bukan ketiganya":
		survey.Details = models.NotWorking{}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid current_activity",
		})
	}

	err := c.ShouldBindJSON(&survey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	createdSurvey, err := h.service.CreateSurvey(survey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdSurvey,
	})
}

func (h *SurveyHandlerImpl) UpdateSurvey(c *gin.Context) {
	survey := models.Survey{}
	err := c.ShouldBindJSON(&survey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	updatedSurvey, err := h.service.UpdateSurvey(survey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedSurvey,
	})
}

func (h *SurveyHandlerImpl) DeleteSurvey(c *gin.Context) {
	deletedSurvey, err := h.service.DeleteSurvey(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": deletedSurvey,
	})
}

func (h *SurveyHandlerImpl) PostSurveyJson(c *gin.Context) {
	var survey map[string]interface{}

	// Bind JSON to the survey map
	err := c.ShouldBindJSON(&survey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	createdSurvey, err := h.service.PostJsonSurvey(survey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdSurvey,
	})
}

func (h *SurveyHandlerImpl) PostSurveyJsons(c *gin.Context) {
	var surveys []map[string]interface{}

	// Bind JSON to the surveys slice
	err := c.ShouldBindJSON(&surveys)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	createdSurveys, err := h.service.PostJsonSurveys(surveys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdSurveys,
	})
}

func (h *SurveyHandlerImpl) GetFeedbacks(c *gin.Context) {
	feedbacks, err := h.service.GetFeedbacks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Printf("Received feedbacks: %v", feedbacks)

	// Process feedbacks with Gemini
	processedFeedbacks, err := ProcessWithGemini(feedbacks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error processing feedbacks: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": processedFeedbacks,
	})
}

func ProcessWithGemini(feedbacks []string) ([]string, error) {
	location := "us-central1"
	modelName := "gemini-1.5-flash-001"

	projectID := configs.EnvGCloudProjectID()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	defer client.Close()

	gemini := client.GenerativeModel(modelName)
	gemini.SetTemperature(0.9)
	gemini.SetTopP(0.5)
	gemini.SetTopK(20)
	gemini.SetMaxOutputTokens(500)

	var processedFeedbacks []string
	for _, feedback := range feedbacks {
		prompt := fmt.Sprintf(`
		Analyze the following feedback and provide suggestions for improvement:
		
		Feedback: %s
		
		Instructions:
		1. Identify the main points of the feedback.
		2. Suggest specific improvements based on the feedback.
		3. Provide any additional insights or recommendations.
		4. Summarize the key takeaways.
		
		Please format your response in a clear and concise manner.
		`, feedback)

		resp, err := gemini.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			return nil, fmt.Errorf("error generating content: %w", err)
		}

		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return nil, fmt.Errorf("no content generated for feedback: %s", feedback)
		}

		// Extract the text content from the response
		content, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
		if !ok {
			return nil, fmt.Errorf("unexpected content type for feedback: %s", feedback)
		}

		processedFeedbacks = append(processedFeedbacks, string(content))
	}

	return processedFeedbacks, nil
}
