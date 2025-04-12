package models

import (
	"time"

	"gorm.io/gorm"
)

// QuizCategory tells us what subject the quiz is about
// For example: math, science, history, etc.
type QuizCategory string

// These are all the different quiz categories we support
const (
	CategoryGeneral    QuizCategory = "general"
	CategoryScience    QuizCategory = "science"
	CategoryHistory    QuizCategory = "history"
	CategoryTechnology QuizCategory = "technology"
	CategoryMath       QuizCategory = "math"
	CategoryLanguage   QuizCategory = "language"
)

// QuizDifficulty tells us how hard the quiz is
type QuizDifficulty string

// These are the three difficulty levels
const (
	DifficultyEasy   QuizDifficulty = "easy"
	DifficultyMedium QuizDifficulty = "medium"
	DifficultyHard   QuizDifficulty = "hard"
)

// QuestionType defines the type of question
type QuestionType string

const (
	QuestionTypeMultipleChoice QuestionType = "multiple_choice"
	QuestionTypeTrueFalse      QuestionType = "true_false"
	QuestionTypeShortAnswer    QuestionType = "short_answer"
	QuestionTypeEssay          QuestionType = "essay"
)

// Quiz represents a complete quiz that users can take
// It includes all the quiz information and its questions
type Quiz struct {
	gorm.Model
	Title        string         `json:"title" binding:"required"`                       // Name of the quiz
	Description  string         `json:"description" binding:"required"`                 // What the quiz is about
	Category     QuizCategory   `json:"category" binding:"required"`                    // Subject area (math, science, etc.)
	Difficulty   QuizDifficulty `json:"difficulty" binding:"required"`                  // How hard it is (easy, medium, hard)
	TimeLimit    int            `json:"time_limit" binding:"required,min=1"`            // How many minutes users have
	Questions    []Question     `json:"questions" binding:"required,min=1"`             // The actual quiz questions
	CreatedBy    uint           `json:"created_by" binding:"required"`                  // Who made the quiz
	IsPublished  bool           `json:"is_published" gorm:"default:false"`              // Whether it's ready for users
	CreatedAt    time.Time      `json:"created_at"`                                     // When it was made
	UpdatedAt    time.Time      `json:"updated_at"`                                     // When it was last changed
	PassingScore float64        `json:"passing_score" binding:"required,min=0,max=100"` // Score needed to pass
}

// Question represents a single question in a quiz
type Question struct {
	gorm.Model
	QuizID        uint         `json:"quiz_id" binding:"required"`              // Which quiz this belongs to
	Text          string       `json:"text" binding:"required"`                 // The actual question
	Type          QuestionType `json:"type" binding:"required"`                 // What kind of question (multiple choice, etc.)
	Answers       []Answer     `json:"answers" binding:"required,min=2"`        // Possible answers
	CorrectAnswer string       `json:"correct_answer" binding:"required"`       // The right answer
	Points        int          `json:"points" gorm:"default:1" binding:"min=1"` // How many points it's worth
	Explanation   string       `json:"explanation"`                             // Why the answer is correct
	TimeToAnswer  int          `json:"time_to_answer" binding:"min=0"`          // Seconds allowed for this question
}

// Answer represents one possible answer to a question
type Answer struct {
	gorm.Model
	QuestionID uint   `json:"question_id" binding:"required"` // Which question this answers
	Text       string `json:"text" binding:"required"`        // The answer text
	IsCorrect  bool   `json:"is_correct"`                     // Whether this is the right answer
	Order      int    `json:"order" binding:"min=0"`          // Display order of the answer
}

// Result stores how a user did on a quiz
type Result struct {
	gorm.Model
	QuizID         uint    `json:"quiz_id" binding:"required"`                     // Which quiz they took
	UserID         uint    `json:"user_id" binding:"required"`                     // Who took the quiz
	Score          float64 `json:"score" binding:"required,min=0,max=100"`         // Their score (0-100)
	TotalQuestions int     `json:"total_questions" binding:"required,min=1"`       // How many questions
	CorrectAnswers int     `json:"correct_answers" binding:"required,min=0"`       // How many they got right
	TimeTaken      int     `json:"time_taken" binding:"required,min=0"`            // How long it took (seconds)
	Answers        []byte  `json:"answers" binding:"required"`                     // Their answers (stored as JSON)
	Feedback       string  `json:"feedback"`                                       // Any feedback for the user
	IsPassed       bool    `json:"is_passed"`                                      // Whether they passed
	PassingScore   float64 `json:"passing_score" binding:"required,min=0,max=100"` // Score needed to pass
}

// UserProgress tracks how well a user is doing in a subject
type UserProgress struct {
	gorm.Model
	UserID          uint         `json:"user_id" binding:"required"`            // Who this progress is for
	QuizID          uint         `json:"quiz_id" binding:"required"`            // Which quiz they took
	Category        QuizCategory `json:"category" binding:"required"`           // In what subject
	TotalAttempts   int          `json:"total_attempts" binding:"min=0"`        // How many times they tried
	BestScore       float64      `json:"best_score" binding:"min=0,max=100"`    // Their highest score
	AverageScore    float64      `json:"average_score" binding:"min=0,max=100"` // Their average score
	TotalTimeSpent  int          `json:"total_time_spent" binding:"min=0"`      // Total time spent (seconds)
	LastAttemptedAt time.Time    `json:"last_attempted_at"`                     // When they last tried
	MasteryLevel    int          `json:"mastery_level" binding:"min=1,max=5"`   // How well they know it (1-5)
}

// IsValidQuizCategory checks if a category name is one we support
func IsValidQuizCategory(category string) bool {
	validCategories := []QuizCategory{
		CategoryGeneral,
		CategoryScience,
		CategoryHistory,
		CategoryTechnology,
		CategoryMath,
		CategoryLanguage,
	}

	for _, c := range validCategories {
		if string(c) == category {
			return true
		}
	}
	return false
}

// IsValidQuizDifficulty checks if a difficulty level is valid
func IsValidQuizDifficulty(difficulty string) bool {
	validDifficulties := []QuizDifficulty{
		DifficultyEasy,
		DifficultyMedium,
		DifficultyHard,
	}

	for _, d := range validDifficulties {
		if string(d) == difficulty {
			return true
		}
	}
	return false
}

// IsValidQuestionType checks if a question type is valid
func IsValidQuestionType(questionType string) bool {
	validTypes := []QuestionType{
		QuestionTypeMultipleChoice,
		QuestionTypeTrueFalse,
		QuestionTypeShortAnswer,
		QuestionTypeEssay,
	}

	for _, t := range validTypes {
		if string(t) == questionType {
			return true
		}
	}
	return false
}

// Question represents a quiz question
type Question struct {
	ID       int64     `json:"id"`
	Text     string    `json:"text"`
	Category string    `json:"category"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// Quiz represents a complete quiz
type Quiz struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
	Created   time.Time  `json:"created_at"`
	Updated   time.Time  `json:"updated_at"`
}

// QuizResponse represents the API response for quiz operations
type QuizResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    *Quiz  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// QuestionResponse represents the API response for question operations
type QuestionResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    *Question `json:"data,omitempty"`
	Error   string    `json:"error,omitempty"`
}
