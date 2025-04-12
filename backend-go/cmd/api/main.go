package main

import (
	"fmt"
	"log"
	"os"

	"aicg/internal/config"
	"aicg/internal/database"
	"aicg/internal/handlers"
	"aicg/internal/middleware"
	"aicg/internal/routes"
	"aicg/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			return
		}
		sqlDB.Close()
	}()

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize services
	authService := services.NewAuthService(db, cfg)
	quizService := services.NewQuizService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	quizHandler := handlers.NewQuizHandler(quizService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg)

	// Initialize Gin router
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	// Public routes
	public := r.Group("/api")
	{
		// Auth routes
		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.GET("/sso/:provider/config", authHandler.GetSSOConfig)
			auth.POST("/sso/:provider/callback", authHandler.HandleSSO)
		}
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(authMiddleware.AuthRequired())
	{
		// Quiz routes
		quiz := protected.Group("/quiz")
		{
			quiz.GET("/", quizHandler.GetQuizzes)
			quiz.GET("/:id", quizHandler.GetQuiz)
			quiz.POST("/:id/submit", quizHandler.SubmitQuiz)
		}

		// Results routes
		results := protected.Group("/results")
		{
			results.GET("/", quizHandler.GetResults)
			results.GET("/:id", quizHandler.GetResult)
		}

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(authMiddleware.RequireSuperAdmin())
		{
			admin.POST("/quiz", quizHandler.CreateQuiz)
			// Add more admin routes here
		}
	}

	// Setup routes
	routes.SetupRoutes(r, quizService)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
