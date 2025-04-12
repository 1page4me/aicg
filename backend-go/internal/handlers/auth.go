package handlers

import (
	"net/http"

	"aicg/internal/models"
	"aicg/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// HandleSSO handles SSO authentication
func (h *AuthHandler) HandleSSO(c *gin.Context) {
	provider := models.AuthProvider(c.Param("provider"))
	if provider != models.ProviderGoogle && provider != models.ProviderFacebook && provider != models.ProviderInstagram {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported SSO provider"})
		return
	}

	var req struct {
		ProviderID string `json:"provider_id" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		FirstName  string `json:"first_name" binding:"required"`
		LastName   string `json:"last_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.authService.HandleSSO(provider, req.ProviderID, req.Email, req.FirstName, req.LastName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// GetSSOConfig returns SSO configuration for a provider
func (h *AuthHandler) GetSSOConfig(c *gin.Context) {
	provider := models.AuthProvider(c.Param("provider"))
	if provider != models.ProviderGoogle && provider != models.ProviderFacebook && provider != models.ProviderInstagram {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported SSO provider"})
		return
	}

	config, err := h.authService.GetSSOConfig(provider)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SSO configuration not found"})
		return
	}

	// Return only necessary information for client
	c.JSON(http.StatusOK, gin.H{
		"client_id":    config.ClientID,
		"redirect_url": config.RedirectURL,
		"scopes":       config.Scopes,
	})
}
