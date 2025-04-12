package routes

import (
	"net/http"

	"aicg/internal/handlers"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all API routes for the application
//
// Parameters:
//   - router: The Gorilla Mux router instance
//
// Returns:
//   - None
func RegisterRoutes(router *mux.Router) {
	// Create API subrouter with /api prefix
	api := router.PathPrefix("/api").Subrouter()

	// Register question routes
	registerQuestionRoutes(api)

	// Register result routes
	registerResultRoutes(api)

	// Register user routes
	registerUserRoutes(api)

	// Register auth routes
	registerAuthRoutes(api)

	// Register health check route
	registerHealthCheck(router)
}

// registerQuestionRoutes sets up routes for question management
func registerQuestionRoutes(api *mux.Router) {
	api.HandleFunc("/questions",
		handlers.GetQuestions).Methods("GET")
	api.HandleFunc("/questions",
		handlers.CreateQuestion).Methods("POST")
}

// registerResultRoutes sets up routes for result management
func registerResultRoutes(api *mux.Router) {
	api.HandleFunc("/results",
		handlers.GetResults).Methods("GET")
	api.HandleFunc("/results",
		handlers.CreateResult).Methods("POST")
}

// registerUserRoutes sets up routes for user management
func registerUserRoutes(api *mux.Router) {
	api.HandleFunc("/users",
		handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}",
		handlers.GetUser).Methods("GET")
	api.HandleFunc("/users",
		handlers.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}",
		handlers.UpdateUser).Methods("PUT")
}

// registerAuthRoutes sets up routes for authentication
func registerAuthRoutes(api *mux.Router) {
	api.HandleFunc("/auth/login",
		handlers.Login).Methods("POST")
	api.HandleFunc("/auth/register",
		handlers.Register).Methods("POST")
	api.HandleFunc("/auth/refresh",
		handlers.RefreshToken).Methods("POST")
}

// registerHealthCheck sets up the health check endpoint
func registerHealthCheck(router *mux.Router) {
	router.HandleFunc("/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}).Methods("GET")
}
