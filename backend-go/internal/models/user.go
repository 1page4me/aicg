package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleSuperAdmin UserRole = "super_admin"
	RoleMaveric    UserRole = "maveric"
)

type AuthProvider string

const (
	ProviderEmail     AuthProvider = "email"
	ProviderGoogle    AuthProvider = "google"
	ProviderFacebook  AuthProvider = "facebook"
	ProviderInstagram AuthProvider = "instagram"
)

type User struct {
	gorm.Model
	Email        string       `json:"email" gorm:"uniqueIndex"`
	PasswordHash string       `json:"-"` // Only for email/password auth
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Role         UserRole     `json:"role" gorm:"default:maveric"`
	AuthProvider AuthProvider `json:"auth_provider" gorm:"default:email"`
	ProviderID   string       `json:"provider_id"` // ID from SSO provider
	LastLoginAt  time.Time    `json:"last_login_at"`
	IsActive     bool         `json:"is_active" gorm:"default:true"`
	ProfileImage string       `json:"profile_image"`
	RefreshToken string       `json:"-"` // For JWT refresh

	// Relationships
	Quizzes          []Quiz            `json:"quizzes" gorm:"foreignKey:CreatedBy"`
	Results          []Result          `json:"results" gorm:"foreignKey:UserID"`
	UserProgress     []UserProgress    `json:"user_progress" gorm:"foreignKey:UserID"`
	UserAchievements []UserAchievement `json:"user_achievements" gorm:"foreignKey:UserID"`
	GlobalRankings   []GlobalRanking   `json:"global_rankings" gorm:"foreignKey:UserID"`
	CategoryRankings []CategoryRanking `json:"category_rankings" gorm:"foreignKey:UserID"`
}

type SSOConfig struct {
	gorm.Model
	Provider     AuthProvider `json:"provider" gorm:"uniqueIndex"`
	ClientID     string       `json:"client_id"`
	ClientSecret string       `json:"client_secret"`
	RedirectURL  string       `json:"redirect_url"`
	Scopes       string       `json:"scopes"` // Comma-separated scopes
	IsEnabled    bool         `json:"is_enabled" gorm:"default:true"`
}
