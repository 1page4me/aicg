package models

// Question represents a quiz question
type Question struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

// QuestionResponse represents the response for question-related operations
type QuestionResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    *Question `json:"data,omitempty"`
}

// Quiz represents a collection of questions
type Quiz struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

// QuizResponse represents the response for quiz-related operations
type QuizResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    *Quiz  `json:"data,omitempty"`
}
