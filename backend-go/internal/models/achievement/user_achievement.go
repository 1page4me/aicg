package achievement

import (
	"time"

	"github.com/yourusername/yourproject/internal/models/user"
	"gorm.io/gorm"
)

// UserAchievement represents an achievement earned by a user
type UserAchievement struct {
	gorm.Model
	UserID        uint        `json:"user_id" gorm:"not null;uniqueIndex:idx_user_achievement"`
	AchievementID uint        `json:"achievement_id" gorm:"not null;uniqueIndex:idx_user_achievement"`
	User          user.User   `json:"user" gorm:"foreignKey:UserID"`
	Achievement   Achievement `json:"achievement" gorm:"foreignKey:AchievementID"`
	EarnedAt      time.Time   `json:"earned_at" gorm:"default:CURRENT_TIMESTAMP"`
	Progress      []byte      `json:"progress" gorm:"type:jsonb"` // Store progress towards achievement
}

// TableName specifies the table name for the UserAchievement model
func (UserAchievement) TableName() string {
	return "user_achievements"
}
