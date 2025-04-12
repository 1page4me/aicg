package handlers

import (
	"aicg/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuiz(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/quizzes/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuiz)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var response models.QuizResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.Data == nil {
		t.Error("Expected data to be present")
	}

	if len(response.Data.Questions) == 0 {
		t.Error("Expected questions to be present")
	}
}

func TestCreateQuiz(t *testing.T) {
	quiz := models.Quiz{
		Title: "Test Quiz",
		Questions: []models.Question{
			{
				Text:     "Test question",
				Category: "Test",
			},
		},
	}

	jsonData, err := json.Marshal(quiz)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(
		"POST",
		"/api/quizzes",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateQuiz)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var response models.QuizResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.Data == nil {
		t.Error("Expected data to be present")
	}

	if response.Data.Title != quiz.Title {
		t.Errorf(
			"Expected title %v, got %v",
			quiz.Title,
			response.Data.Title,
		)
	}
}
