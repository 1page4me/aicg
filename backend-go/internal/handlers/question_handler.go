package handlers

import (
	"aicg/internal/models"
	"encoding/json"
	"net/http"
)

// GetQuestions handles GET requests to fetch all questions
//
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
//
// Returns:
//   - None
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement database query
	questions := []models.Question{
		{
			ID:       1,
			Text:     "Sample question 1",
			Category: "General",
		},
	}

	response := models.QuestionResponse{
		Success: true,
		Data:    &questions[0],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateQuestion handles POST requests to create a new question
//
// Parameters:
//   - w: HTTP response writer
//   - r: HTTP request
//
// Returns:
//   - None
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Implement database insert
	question.ID = 1

	response := models.QuestionResponse{
		Success: true,
		Message: "Question created successfully",
		Data:    &question,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
