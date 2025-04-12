package ranking

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"github.com/yourusername/yourproject/internal/models/user"
	"gorm.io/gorm"
)

// GlobalRanking represents a user's overall performance across all categories
type GlobalRanking struct {
	gorm.Model
	UserID           uint                `json:"user_id" gorm:"not null;uniqueIndex:idx_user_period"`
	User             user.User           `json:"user" gorm:"foreignKey:UserID"`
	TotalScore       float64             `json:"total_score" gorm:"type:decimal(10,2);default:0"`
	QuizzesCompleted int                 `json:"quizzes_completed" gorm:"default:0"`
	AverageScore     float64             `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	TotalTimeSpent   int                 `json:"total_time_spent" gorm:"default:0"` // in seconds
	Rank             int                 `json:"rank"`
	Percentile       float64             `json:"percentile" gorm:"type:decimal(5,2)"`
	RankingPeriod    enums.RankingPeriod `json:"ranking_period" gorm:"size:20;not null;uniqueIndex:idx_user_period"`
}

// TableName specifies the table name for the GlobalRanking model
func (GlobalRanking) TableName() string {
	return "global_rankings"
}
