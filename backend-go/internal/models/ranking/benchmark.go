package ranking

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// Benchmark represents performance metrics for a specific category and difficulty
type Benchmark struct {
	gorm.Model
	Category              enums.QuizCategory   `json:"category" gorm:"size:50;not null;uniqueIndex:idx_category_difficulty"`
	Difficulty            enums.QuizDifficulty `json:"difficulty" gorm:"size:20;not null;uniqueIndex:idx_category_difficulty"`
	AverageScore          float64              `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	MedianScore           float64              `json:"median_score" gorm:"type:decimal(5,2);default:0"`
	Percentile75          float64              `json:"percentile_75" gorm:"type:decimal(5,2);default:0"`
	Percentile90          float64              `json:"percentile_90" gorm:"type:decimal(5,2);default:0"`
	TotalAttempts         int                  `json:"total_attempts" gorm:"default:0"`
	AverageCompletionTime int                  `json:"average_completion_time" gorm:"default:0"` // in seconds
}

// TableName specifies the table name for the Benchmark model
func (Benchmark) TableName() string {
	return "benchmarks"
}
