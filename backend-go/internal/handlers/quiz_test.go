package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"aicg/internal/models"
	"aicg/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuizService is a mock implementation of IQuizService
type MockQuizService struct {
	mock.Mock
}

// Ensure MockQuizService implements services.IQuizService
var _ services.IQuizService = (*MockQuizService)(nil)

func (m *MockQuizService) GetQuizzes() ([]models.Quiz, error) {
	args := m.Called()
	return args.Get(0).([]models.Quiz), args.Error(1)
}

func (m *MockQuizService) GetQuizByID(id uint) (*models.Quiz, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Quiz), args.Error(1)
}

func (m *MockQuizService) CreateQuiz(quiz *models.Quiz) error {
	args := m.Called(quiz)
	return args.Error(0)
}

func (m *MockQuizService) GetQuizzesByCategory(category models.QuizCategory) ([]models.Quiz, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Quiz), args.Error(1)
}

func (m *MockQuizService) GetQuizzesByDifficulty(difficulty models.QuizDifficulty) ([]models.Quiz, error) {
	args := m.Called(difficulty)
	return args.Get(0).([]models.Quiz), args.Error(1)
}

func (m *MockQuizService) SubmitQuizResult(result *models.Result) error {
	args := m.Called(result)
	return args.Error(0)
}

func (m *MockQuizService) GetUserResults(userID uint) ([]models.Result, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Result), args.Error(1)
}

func (m *MockQuizService) GetResultByID(id uint) (*models.Result, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Result), args.Error(1)
}

func (m *MockQuizService) GetUserProgress(userID uint) ([]models.UserProgress, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.UserProgress), args.Error(1)
}

func (m *MockQuizService) GetUserProgressByCategory(userID uint, category models.QuizCategory) (*models.UserProgress, error) {
	args := m.Called(userID, category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserProgress), args.Error(1)
}

func setupTest() (*gin.Engine, *MockQuizService) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockService := new(MockQuizService)
	handler := NewQuizHandler(mockService)

	// Setup routes
	r.GET("/quizzes", handler.GetQuizzes)
	r.GET("/quiz/:id", handler.GetQuiz)
	r.POST("/quiz", handler.CreateQuiz)
	r.GET("/quiz/category/:category", handler.GetQuizzesByCategory)
	r.GET("/quiz/difficulty/:difficulty", handler.GetQuizzesByDifficulty)
	r.POST("/quiz/:id/submit", handler.SubmitQuiz)
	r.GET("/results", handler.GetResults)
	r.GET("/result/:id", handler.GetResult)

	return r, mockService
}

func TestGetQuizzes(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful retrieval
	quizzes := []models.Quiz{{Title: "Test Quiz"}}
	mockService.On("GetQuizzes").Return(quizzes, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quizzes", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Quiz
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, quizzes, response)
}

func TestGetQuiz(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful retrieval
	quiz := &models.Quiz{Title: "Test Quiz"}
	mockService.On("GetQuizByID", uint(1)).Return(quiz, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quiz/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Quiz
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, quiz.Title, response.Title)

	// Test case 2: Invalid ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/quiz/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateQuiz(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful creation
	quiz := &models.Quiz{Title: "New Quiz"}
	mockService.On("CreateQuiz", mock.AnythingOfType("*models.Quiz")).Return(nil)

	body, _ := json.Marshal(quiz)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/quiz", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetQuizzesByCategory(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful retrieval
	quizzes := []models.Quiz{{Title: "Math Quiz"}}
	mockService.On("GetQuizzesByCategory", models.CategoryMath).Return(quizzes, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/quiz/category/math", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Quiz
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, quizzes, response)

	// Test case 2: Invalid category
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/quiz/category/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetResults(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful retrieval
	results := []models.Result{{Score: 85}}
	mockService.On("GetUserResults", uint(1)).Return(results, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/results?user_id=1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, results, response)

	// Test case 2: Invalid user ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/results?user_id=invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetResult(t *testing.T) {
	r, mockService := setupTest()

	// Test case 1: Successful retrieval
	result := &models.Result{Score: 85}
	mockService.On("GetResultByID", uint(1)).Return(result, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/result/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Result
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, result.Score, response.Score)

	// Test case 2: Invalid ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/result/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
