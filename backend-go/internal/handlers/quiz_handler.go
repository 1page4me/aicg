package handlers

import (
	"net/http"

	"aicg/backend-go/internal/models"

	"github.com/gin-gonic/gin"
)

// GetQuiz handles GET requests to fetch a quiz
func GetQuiz(c *gin.Context) {
	// Sample quiz for testing
	sampleQuiz := &models.Quiz{
		ID:    1,
		Title: "Sample Quiz",
		Questions: []models.Question{
			{
				ID:       1,
				Text:     "What is the capital of France?",
				Category: "Geography",
			},
			{
				ID:       2,
				Text:     "What is 2 + 2?",
				Category: "Math",
			},
		},
	}

	response := models.QuizResponse{
		Success: true,
		Message: "Quiz retrieved successfully",
		Data:    sampleQuiz,
	}

	c.JSON(http.StatusOK, response)
}

// CreateQuiz handles POST requests to create a new quiz
func CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		response := models.QuizResponse{
			Success: false,
			Message: "Invalid request body",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Here you would typically save the quiz to a database
	// For now, we'll just return the received quiz

	response := models.QuizResponse{
		Success: true,
		Message: "Quiz created successfully",
		Data:    &quiz,
	}

	c.JSON(http.StatusCreated, response)
}
