package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterRoutes(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()

	// Register routes
	RegisterRoutes(router)

	// Test cases
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{
			name:       "Health check",
			method:     "GET",
			path:       "/health",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get questions",
			method:     "GET",
			path:       "/api/questions",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Create question",
			method:     "POST",
			path:       "/api/questions",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get results",
			method:     "GET",
			path:       "/api/results",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Create result",
			method:     "POST",
			path:       "/api/results",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get users",
			method:     "GET",
			path:       "/api/users",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get user by ID",
			method:     "GET",
			path:       "/api/users/1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Create user",
			method:     "POST",
			path:       "/api/users",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Update user",
			method:     "PUT",
			path:       "/api/users/1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Login",
			method:     "POST",
			path:       "/api/auth/login",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Register",
			method:     "POST",
			path:       "/api/auth/register",
			wantStatus: http.StatusOK,
		},
		{
			name:       "Refresh token",
			method:     "POST",
			path:       "/api/auth/refresh",
			wantStatus: http.StatusOK,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.wantStatus,
				)
			}
		})
	}
}
