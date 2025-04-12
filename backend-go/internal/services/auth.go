package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"aicg/internal/config"
	"aicg/internal/models"
)

type AuthService struct {
	db     *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		db:     db,
		config: cfg,
	}
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Register a new user
func (s *AuthService) Register(email, password, firstName, lastName string) (*models.User, error) {
	// Check if user already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		FirstName:    firstName,
		LastName:     lastName,
		Role:         models.RoleMaveric,
		AuthProvider: models.ProviderEmail,
		IsActive:     true,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Login with email/password
func (s *AuthService) Login(email, password string) (*TokenPair, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Update last login
	user.LastLoginAt = time.Now()
	s.db.Save(&user)

	return s.generateTokens(&user)
}

// Handle SSO login/registration
func (s *AuthService) HandleSSO(provider models.AuthProvider, providerID, email, firstName, lastName string) (*TokenPair, error) {
	var user models.User
	err := s.db.Where("provider_id = ? AND auth_provider = ?", providerID, provider).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		// Create new user
		user = models.User{
			Email:        email,
			FirstName:    firstName,
			LastName:     lastName,
			Role:         models.RoleMaveric,
			AuthProvider: provider,
			ProviderID:   providerID,
			IsActive:     true,
		}
		if err := s.db.Create(&user).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Update last login
	user.LastLoginAt = time.Now()
	s.db.Save(&user)

	return s.generateTokens(&user)
}

// Generate JWT tokens
func (s *AuthService) generateTokens(user *models.User) (*TokenPair, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(s.config.JWTExpiry).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	// Save refresh token
	user.RefreshToken = refreshTokenString
	s.db.Save(user)

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// Refresh access token
func (s *AuthService) RefreshToken(refreshToken string) (*TokenPair, error) {
	var user models.User
	if err := s.db.Where("refresh_token = ?", refreshToken).First(&user).Error; err != nil {
		return nil, errors.New("invalid refresh token")
	}

	return s.generateTokens(&user)
}

// Get SSO configuration
func (s *AuthService) GetSSOConfig(provider models.AuthProvider) (*models.SSOConfig, error) {
	var config models.SSOConfig
	if err := s.db.Where("provider = ? AND is_enabled = ?", provider, true).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}
