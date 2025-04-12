package achievement

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// Achievement represents an achievement that users can earn
type Achievement struct {
	gorm.Model
	Name             string             `json:"name" gorm:"size:100;not null"`
	Description      string             `json:"description" gorm:"type:text;not null"`
	Category         enums.QuizCategory `json:"category" gorm:"size:50"`
	Criteria         []byte             `json:"criteria" gorm:"type:jsonb;not null"` // Store achievement criteria
	IconURL          string             `json:"icon_url" gorm:"size:255"`
	UserAchievements []UserAchievement  `json:"user_achievements" gorm:"foreignKey:AchievementID"`
}

// TableName specifies the table name for the Achievement model
func (Achievement) TableName() string {
	return "achievements"
}
