package main

import (
	"log"
	"net/http"
	"os"

	"aicg/internal/middleware"
	"aicg/internal/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Apply CORS middleware
	router.Use(middleware.CORS)

	// Register routes
	routes.RegisterRoutes(router)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
