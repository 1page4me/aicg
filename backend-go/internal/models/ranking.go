package models

import (
	"gorm.io/gorm"
)

type RankingPeriod string

const (
	RankingPeriodWeekly  RankingPeriod = "weekly"
	RankingPeriodMonthly RankingPeriod = "monthly"
	RankingPeriodAllTime RankingPeriod = "all_time"
)

// GlobalRanking represents a user's overall performance across all categories
type GlobalRanking struct {
	gorm.Model
	UserID           uint          `json:"user_id" gorm:"not null;uniqueIndex:idx_user_period"`
	User             User          `json:"user" gorm:"foreignKey:UserID"`
	TotalScore       float64       `json:"total_score" gorm:"type:decimal(10,2);default:0"`
	QuizzesCompleted int           `json:"quizzes_completed" gorm:"default:0"`
	AverageScore     float64       `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	TotalTimeSpent   int           `json:"total_time_spent" gorm:"default:0"` // in seconds
	Rank             int           `json:"rank"`
	Percentile       float64       `json:"percentile" gorm:"type:decimal(5,2)"`
	RankingPeriod    RankingPeriod `json:"ranking_period" gorm:"size:20;not null;uniqueIndex:idx_user_period"`
}

// CategoryRanking represents a user's performance in a specific category
type CategoryRanking struct {
	gorm.Model
	UserID           uint          `json:"user_id" gorm:"not null;uniqueIndex:idx_user_category_period"`
	User             User          `json:"user" gorm:"foreignKey:UserID"`
	Category         QuizCategory  `json:"category" gorm:"size:50;not null;uniqueIndex:idx_user_category_period"`
	TotalScore       float64       `json:"total_score" gorm:"type:decimal(10,2);default:0"`
	QuizzesCompleted int           `json:"quizzes_completed" gorm:"default:0"`
	AverageScore     float64       `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	Rank             int           `json:"rank"`
	Percentile       float64       `json:"percentile" gorm:"type:decimal(5,2)"`
	RankingPeriod    RankingPeriod `json:"ranking_period" gorm:"size:20;not null;uniqueIndex:idx_user_category_period"`
}

// Benchmark represents performance metrics for a specific category and difficulty
type Benchmark struct {
	gorm.Model
	Category              QuizCategory   `json:"category" gorm:"size:50;not null;uniqueIndex:idx_category_difficulty"`
	Difficulty            QuizDifficulty `json:"difficulty" gorm:"size:20;not null;uniqueIndex:idx_category_difficulty"`
	AverageScore          float64        `json:"average_score" gorm:"type:decimal(5,2);default:0"`
	MedianScore           float64        `json:"median_score" gorm:"type:decimal(5,2);default:0"`
	Percentile75          float64        `json:"percentile_75" gorm:"type:decimal(5,2);default:0"`
	Percentile90          float64        `json:"percentile_90" gorm:"type:decimal(5,2);default:0"`
	TotalAttempts         int            `json:"total_attempts" gorm:"default:0"`
	AverageCompletionTime int            `json:"average_completion_time" gorm:"default:0"` // in seconds
}

// LeaderboardEntry represents a user's position in the leaderboard
type LeaderboardEntry struct {
	UserID           uint   `json:"user_id"`
	Username         string `json:"username"`
	Score            int    `json:"score"`
	Rank             int    `json:"rank"`
	Category         string `json:"category"`
	AchievementCount int    `json:"achievement_count"`
}
