package middleware

import (
	"net/http"
	"strings"

	"aicg/internal/config"
	"aicg/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware checks if users are logged in and have permission
// to access different parts of the application
type AuthMiddleware struct {
	config *config.Config
}

// NewAuthMiddleware creates a new auth middleware with the app config
func NewAuthMiddleware(cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{config: cfg}
}

// AuthRequired makes sure the user is logged in before accessing a route
// It checks for a valid JWT token in the Authorization header
func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Look for the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Make sure it's a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Check if the token is valid
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get the user information from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Save user info so other parts of the app can use it
		c.Set("userID", uint(claims["sub"].(float64)))
		c.Set("userEmail", claims["email"].(string))
		c.Set("userRole", models.UserRole(claims["role"].(string)))

		c.Next()
	}
}

// RequireRole checks if the user has one of the required roles
// If not, they can't access the route
func (m *AuthMiddleware) RequireRole(roles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user's role from the context
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		// Check if the user's role matches any of the required roles
		role := userRole.(models.UserRole)
		for _, requiredRole := range roles {
			if role == requiredRole {
				c.Next()
				return
			}
		}

		// If we get here, the user doesn't have permission
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		c.Abort()
	}
}

// RequireSuperAdmin is a shortcut to check if user is a super admin
// Only super admins can access routes with this middleware
func (m *AuthMiddleware) RequireSuperAdmin() gin.HandlerFunc {
	return m.RequireRole(models.RoleSuperAdmin)
}
