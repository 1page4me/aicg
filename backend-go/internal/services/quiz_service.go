package services

import "aicg/internal/models"

// IQuizService defines the interface for quiz-related operations
type IQuizService interface {
	// GetQuizzes retrieves all quizzes
	GetQuizzes() ([]models.Quiz, error)

	// GetQuizByID retrieves a specific quiz by its ID
	GetQuizByID(id uint) (*models.Quiz, error)

	// CreateQuiz creates a new quiz
	CreateQuiz(quiz *models.Quiz) error

	// GetQuizzesByCategory retrieves quizzes filtered by category
	GetQuizzesByCategory(category models.QuizCategory) ([]models.Quiz, error)

	// GetQuizzesByDifficulty retrieves quizzes filtered by difficulty
	GetQuizzesByDifficulty(difficulty models.QuizDifficulty) ([]models.Quiz, error)

	// SubmitQuizResult saves a quiz result
	SubmitQuizResult(result *models.Result) error

	// GetUserResults retrieves all quiz results for a user
	GetUserResults(userID uint) ([]models.Result, error)

	// GetResultByID retrieves a specific quiz result by its ID
	GetResultByID(id uint) (*models.Result, error)

	// GetUserProgress retrieves a user's progress across all quizzes
	GetUserProgress(userID uint) ([]models.UserProgress, error)

	// GetUserProgressByCategory retrieves a user's progress for a specific category
	GetUserProgressByCategory(userID uint, category models.QuizCategory) (*models.UserProgress, error)
}
