package user

import (
	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// SSOConfig represents the configuration for a Single Sign-On provider
type SSOConfig struct {
	gorm.Model
	Provider     enums.AuthProvider `json:"provider" gorm:"size:20;uniqueIndex"`
	ClientID     string             `json:"client_id" gorm:"size:255;not null"`
	ClientSecret string             `json:"client_secret" gorm:"size:255;not null"`
	RedirectURL  string             `json:"redirect_url" gorm:"size:255;not null"`
	Scopes       string             `json:"scopes" gorm:"size:255"` // Comma-separated scopes
	IsEnabled    bool               `json:"is_enabled" gorm:"default:true"`
}

// TableName specifies the table name for the SSOConfig model
func (SSOConfig) TableName() string {
	return "sso_configs"
}
