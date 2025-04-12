package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"aicg/internal/database"
	"aicg/internal/models"
	"aicg/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// QuizHandler manages all quiz-related HTTP requests
// It uses a quiz service to handle business logic
type QuizHandler struct {
	quizService services.IQuizService
}

// NewQuizHandler creates a new quiz handler with the given service
func NewQuizHandler(quizService services.IQuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

// GetQuizzes returns a list of all available quizzes
// GET /api/quizzes
func (h *QuizHandler) GetQuizzes(c *gin.Context) {
	quizzes, err := h.quizService.GetQuizzes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

// GetQuiz returns a specific quiz by its ID
// GET /api/quiz/:id
func (h *QuizHandler) GetQuiz(c *gin.Context) {
	// Convert the ID from string to number
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}

	quiz, err := h.quizService.GetQuizByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// CreateQuiz creates a new quiz
// POST /api/quiz
func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	// Try to parse the request body into a Quiz object
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.quizService.CreateQuiz(&quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}

// GetQuizzesByCategory returns all quizzes in a specific category
// GET /api/quiz/category/:category
func (h *QuizHandler) GetQuizzesByCategory(c *gin.Context) {
	category := c.Param("category")
	// Make sure the category is valid
	if !models.IsValidQuizCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	quizzes, err := h.quizService.GetQuizzesByCategory(models.QuizCategory(category))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

// GetQuizzesByDifficulty returns all quizzes of a specific difficulty level
// GET /api/quiz/difficulty/:difficulty
func (h *QuizHandler) GetQuizzesByDifficulty(c *gin.Context) {
	difficulty := c.Param("difficulty")
	// Make sure the difficulty level is valid
	if !models.IsValidQuizDifficulty(difficulty) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid difficulty"})
		return
	}

	quizzes, err := h.quizService.GetQuizzesByDifficulty(models.QuizDifficulty(difficulty))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

// SubmitQuiz handles a user submitting their quiz answers
// POST /api/quiz/:id/submit
func (h *QuizHandler) SubmitQuiz(c *gin.Context) {
	// Get user ID from the authentication context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get quiz ID from URL
	id := c.Param("id")
	quizID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID format"})
		return
	}

	// Structure for the submitted answers
	var submission struct {
		Answers []struct {
			QuestionID uint   `json:"question_id" binding:"required"`
			Answer     string `json:"answer" binding:"required"`
		} `json:"answers" binding:"required,min=1"`
		TimeTaken int `json:"time_taken" binding:"required,min=0"`
	}

	// Parse the submission data
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid submission format: " + err.Error()})
		return
	}

	// Get the quiz with its questions
	var quiz models.Quiz
	if err := database.DB.Preload("Questions").First(&quiz, quizID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quiz"})
		}
		return
	}

	// Make sure all submitted question IDs are valid
	questionIDs := make(map[uint]bool)
	for _, q := range quiz.Questions {
		questionIDs[q.ID] = true
	}
	for _, a := range submission.Answers {
		if !questionIDs[a.QuestionID] {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid question ID: %d", a.QuestionID)})
			return
		}
	}

	// Calculate the score
	correctAnswers := 0
	for _, q := range quiz.Questions {
		for _, a := range submission.Answers {
			if q.ID == a.QuestionID && q.CorrectAnswer == a.Answer {
				correctAnswers++
			}
		}
	}

	// Convert to percentage score
	score := float64(correctAnswers) / float64(len(quiz.Questions)) * 100

	// Save the result
	result := models.Result{
		QuizID:         uint(quizID),
		UserID:         userID.(uint),
		Score:          score,
		TotalQuestions: len(quiz.Questions),
		TimeTaken:      submission.TimeTaken,
		IsPassed:       score >= quiz.PassingScore,
		PassingScore:   quiz.PassingScore,
	}

	if err := database.DB.Create(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save quiz result"})
		return
	}

	// Return the results to the user
	c.JSON(http.StatusOK, gin.H{
		"score":           score,
		"total_questions": len(quiz.Questions),
		"correct_answers": correctAnswers,
		"time_taken":      submission.TimeTaken,
		"is_passed":       result.IsPassed,
		"passing_score":   quiz.PassingScore,
	})
}

// SubmitQuizResult handles POST request to submit quiz result
func (h *QuizHandler) SubmitQuizResult(c *gin.Context) {
	var result models.Result
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.quizService.SubmitQuizResult(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUserProgress handles GET request to fetch user's quiz progress
func (h *QuizHandler) GetUserProgress(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	progress, err := h.quizService.GetUserProgress(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// GetUserProgressByCategory handles GET request to fetch user's progress by category
func (h *QuizHandler) GetUserProgressByCategory(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	category := c.Param("category")
	if !models.IsValidQuizCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	progress, err := h.quizService.GetUserProgressByCategory(uint(userID), models.QuizCategory(category))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// GetResults handles GET request to fetch user's quiz results
func (h *QuizHandler) GetResults(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	results, err := h.quizService.GetUserResults(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetResult handles GET request to fetch a specific quiz result
func (h *QuizHandler) GetResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid result ID"})
		return
	}

	result, err := h.quizService.GetResultByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Result not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
