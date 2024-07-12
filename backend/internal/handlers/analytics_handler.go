package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/vertexai/genai"
	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
)

func GenerateContentFromText(w io.Writer, projectID string, req GenerateContentRequest) error {
	location := "us-central1"
	modelName := "gemini-1.5-flash-001"

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	gemini := client.GenerativeModel(modelName)
	gemini.SetTemperature(0.9)
	gemini.SetTopP(0.5)
	gemini.SetTopK(20)
	gemini.SetMaxOutputTokens(500)
	gemini.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("You're a data analyst and education consultant who specializes in improving school curriculums and give advice on how to improve student performance.")},
	}
	prompt := genai.Text(req.Prompt)

	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return fmt.Errorf("error generating content: %w", err)
	}
	// See the JSON response in
	// https://pkg.go.dev/cloud.google.com/go/vertexai/genai#GenerateContentResponse.
	rb, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %w", err)
	}
	fmt.Fprintln(w, string(rb))
	return nil
}

type GenerateContentRequest struct {
	Prompt string `json:"prompt"` // Field to receive user prompt
}

func GenerateAnalytics(c *gin.Context) {
	var request GenerateContentRequest
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to parse request body"})
		return
	}

	projectID := configs.EnvGCloudProjectID()

	var w = c.Writer
	err = GenerateContentFromText(w, projectID, request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
