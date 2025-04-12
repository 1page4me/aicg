package services

import (
	"errors"
	"time"

	"aicg/internal/models"

	"gorm.io/gorm"
)

type QuizService struct {
	db *gorm.DB
}

func NewQuizService(db *gorm.DB) *QuizService {
	return &QuizService{db: db}
}

func (s *QuizService) CreateQuiz(quiz *models.Quiz) error {
	return s.db.Create(quiz).Error
}

func (s *QuizService) GetQuizByID(id uint) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := s.db.Preload("Questions.Answers").First(&quiz, id).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (s *QuizService) GetQuizzesByCategory(category models.QuizCategory) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := s.db.Where("category = ? AND is_published = ?", category, true).Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (s *QuizService) GetQuizzesByDifficulty(difficulty models.QuizDifficulty) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := s.db.Where("difficulty = ? AND is_published = ?", difficulty, true).Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (s *QuizService) SubmitQuizResult(result *models.Result) error {
	// Calculate mastery level based on score
	masteryLevel := calculateMasteryLevel(result.Score)

	// Update or create user progress
	progress := &models.UserProgress{
		UserID:          result.UserID,
		QuizID:          result.QuizID,
		TotalAttempts:   1,
		BestScore:       result.Score,
		AverageScore:    result.Score,
		TotalTimeSpent:  result.TimeTaken,
		LastAttemptedAt: time.Now(),
		MasteryLevel:    masteryLevel,
	}

	// Get quiz category
	var quiz models.Quiz
	if err := s.db.First(&quiz, result.QuizID).Error; err != nil {
		return err
	}
	progress.Category = quiz.Category

	// Use transaction to ensure data consistency
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Save result
		if err := tx.Create(result).Error; err != nil {
			return err
		}

		// Update or create progress
		var existingProgress models.UserProgress
		if err := tx.Where("user_id = ? AND quiz_id = ?", result.UserID, result.QuizID).First(&existingProgress).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(progress).Error
			}
			return err
		}

		// Update existing progress
		existingProgress.TotalAttempts++
		existingProgress.TotalTimeSpent += result.TimeTaken
		existingProgress.AverageScore = (existingProgress.AverageScore*float64(existingProgress.TotalAttempts-1) + result.Score) / float64(existingProgress.TotalAttempts)
		if result.Score > existingProgress.BestScore {
			existingProgress.BestScore = result.Score
		}
		existingProgress.LastAttemptedAt = time.Now()
		existingProgress.MasteryLevel = calculateMasteryLevel(existingProgress.AverageScore)

		return tx.Save(&existingProgress).Error
	})
}

func (s *QuizService) GetUserProgress(userID uint) ([]models.UserProgress, error) {
	var progress []models.UserProgress
	if err := s.db.Where("user_id = ?", userID).Find(&progress).Error; err != nil {
		return nil, err
	}
	return progress, nil
}

func (s *QuizService) GetUserProgressByCategory(userID uint, category models.QuizCategory) (*models.UserProgress, error) {
	var progress models.UserProgress
	if err := s.db.Where("user_id = ? AND category = ?", userID, category).First(&progress).Error; err != nil {
		return nil, err
	}
	return &progress, nil
}

func (s *QuizService) GetQuizzes() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := s.db.Preload("Questions.Answers").Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func calculateMasteryLevel(score float64) int {
	switch {
	case score >= 90:
		return 5
	case score >= 80:
		return 4
	case score >= 70:
		return 3
	case score >= 60:
		return 2
	default:
		return 1
	}
}

func (s *QuizService) GetUserResults(userID uint) ([]models.Result, error) {
	var results []models.Result
	if err := s.db.Where("user_id = ?", userID).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (s *QuizService) GetResultByID(id uint) (*models.Result, error) {
	var result models.Result
	if err := s.db.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
