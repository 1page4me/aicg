package handlers

import (
	"aicg/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuestions(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetQuestions)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var response models.QuestionResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.Data == nil {
		t.Error("Expected data to be present")
	}
}

func TestCreateQuestion(t *testing.T) {
	question := models.Question{
		Text:     "Test question",
		Category: "Test",
	}

	jsonData, err := json.Marshal(question)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(
		"POST",
		"/api/questions",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateQuestion)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	var response models.QuestionResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Error("Expected success to be true")
	}

	if response.Data == nil {
		t.Error("Expected data to be present")
	}

	if response.Data.Text != question.Text {
		t.Errorf(
			"Expected text %v, got %v",
			question.Text,
			response.Data.Text,
		)
	}
}
