package quiz

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// Quiz represents a complete quiz that users can take
type Quiz struct {
	gorm.Model
	Title        string               `json:"title" gorm:"size:255;not null"`
	Description  string               `json:"description" gorm:"type:text;not null"`
	Category     enums.QuizCategory   `json:"category" gorm:"size:50;not null"`
	Difficulty   enums.QuizDifficulty `json:"difficulty" gorm:"size:20;not null"`
	TimeLimit    int                  `json:"time_limit" gorm:"not null"`
	Questions    []Question           `json:"questions" gorm:"foreignKey:QuizID"`
	CreatedBy    uint                 `json:"created_by" gorm:"not null"`
	IsPublished  bool                 `json:"is_published" gorm:"default:false"`
	PassingScore float64              `json:"passing_score" gorm:"type:decimal(5,2);not null;check:passing_score >= 0 AND passing_score <= 100"`
}

// TableName specifies the table name for the Quiz model
func (Quiz) TableName() string {
	return "quizzes"
}
