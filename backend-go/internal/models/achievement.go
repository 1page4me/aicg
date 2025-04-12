package models

import (
	"time"

	"gorm.io/gorm"
)

type Achievement struct {
	gorm.Model
	Name             string            `json:"name" gorm:"size:100;not null"`
	Description      string            `json:"description" gorm:"type:text;not null"`
	Category         QuizCategory      `json:"category" gorm:"size:50"`
	Criteria         []byte            `json:"criteria" gorm:"type:jsonb;not null"` // Store achievement criteria
	IconURL          string            `json:"icon_url" gorm:"size:255"`
	UserAchievements []UserAchievement `json:"user_achievements" gorm:"foreignKey:AchievementID"`
}

type UserAchievement struct {
	gorm.Model
	UserID        uint        `json:"user_id" gorm:"not null;uniqueIndex:idx_user_achievement"`
	AchievementID uint        `json:"achievement_id" gorm:"not null;uniqueIndex:idx_user_achievement"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	Achievement   Achievement `json:"achievement" gorm:"foreignKey:AchievementID"`
	EarnedAt      time.Time   `json:"earned_at" gorm:"default:CURRENT_TIMESTAMP"`
	Progress      []byte      `json:"progress" gorm:"type:jsonb"` // Store progress towards achievement
}
