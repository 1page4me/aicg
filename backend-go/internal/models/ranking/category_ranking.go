package ranking

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"github.com/yourusername/yourproject/internal/models/user"
	"gorm.io/gorm"
)

// CategoryRanking represents a user's performance in a specific category
type CategoryRanking struct {
	gorm.Model
	UserID           uint                `json:"user_id" gorm:"not null;uniqueIndex:idx_user_category_period"`
	User             user.User           `json:"user" gorm:"foreignKey:UserID"`
	Category         enums.QuizCategory  `json:"category" gorm:"size:50;not null;uniqueIndex:idx_user_category_period"`
	TotalScore       float64             `json:"total_score" gorm:"type:decimal(10,2);default:0"`
	QuizzesCompleted int                 `json:"quizzes_completed" gorm:"default:0"`
	AverageScore     float64             `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	Rank             int                 `json:"rank"`
	Percentile       float64             `json:"percentile" gorm:"type:decimal(5,2)"`
	RankingPeriod    enums.RankingPeriod `json:"ranking_period" gorm:"size:20;not null;uniqueIndex:idx_user_category_period"`
}

// TableName specifies the table name for the CategoryRanking model
func (CategoryRanking) TableName() string {
	return "category_rankings"
}
