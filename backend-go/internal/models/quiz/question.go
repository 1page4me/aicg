package quiz

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// Question represents a single question in a quiz
type Question struct {
	gorm.Model
	QuizID        uint               `json:"quiz_id" gorm:"not null"`
	Text          string             `json:"text" gorm:"type:text;not null"`
	Type          enums.QuestionType `json:"type" gorm:"size:20;not null"`
	Answers       []Answer           `json:"answers" gorm:"foreignKey:QuestionID"`
	CorrectAnswer string             `json:"correct_answer" gorm:"type:text;not null"`
	Points        int                `json:"points" gorm:"default:1;check:points >= 1"`
	Explanation   string             `json:"explanation" gorm:"type:text"`
	TimeToAnswer  int                `json:"time_to_answer" gorm:"default:30;check:time_to_answer >= 0"`
}

// TableName specifies the table name for the Question model
func (Question) TableName() string {
	return "questions"
}
