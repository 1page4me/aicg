package user

import (
	"time"

	"github.com/yourusername/yourproject/internal/models/enums"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Email        string             `json:"email" gorm:"size:255;uniqueIndex;not null"`
	PasswordHash string             `json:"-" gorm:"size:255"` // Only for email/password auth
	FirstName    string             `json:"first_name" gorm:"size:100;not null"`
	LastName     string             `json:"last_name" gorm:"size:100;not null"`
	Role         enums.UserRole     `json:"role" gorm:"size:20;not null;default:maveric"`
	AuthProvider enums.AuthProvider `json:"auth_provider" gorm:"size:20;not null;default:email"`
	ProviderID   string             `json:"provider_id" gorm:"size:255"` // ID from SSO provider
	LastLoginAt  time.Time          `json:"last_login_at"`
	IsActive     bool               `json:"is_active" gorm:"default:true"`
	ProfileImage string             `json:"profile_image" gorm:"size:255"`
	RefreshToken string             `json:"-" gorm:"type:text"` // For JWT refresh
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
