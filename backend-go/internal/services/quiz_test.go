package services

import (
	"regexp"
	"testing"

	"aicg/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *QuizService) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM DB: %v", err)
	}

	return db, mock, NewQuizService(db)
}

func TestGetQuizzes(t *testing.T) {
	_, mock, service := setupTestDB(t)

	// Mock quizzes query
	quizRows := sqlmock.NewRows([]string{"id", "title", "description", "deleted_at"}).
		AddRow(1, "Test Quiz", "Test Description", nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quizzes` WHERE `quizzes`.`deleted_at` IS NULL")).
		WillReturnRows(quizRows)

	// Mock questions query
	questionRows := sqlmock.NewRows([]string{"id", "quiz_id", "text", "correct_answer", "deleted_at"}).
		AddRow(1, 1, "Test Question", "A", nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `questions` WHERE `questions`.`quiz_id` = ? AND `questions`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(questionRows)

	// Mock answers query
	answerRows := sqlmock.NewRows([]string{"id", "question_id", "text", "is_correct", "deleted_at"}).
		AddRow(1, 1, "Answer A", true, nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `answers` WHERE `answers`.`question_id` = ? AND `answers`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(answerRows)

	quizzes, err := service.GetQuizzes()
	assert.NoError(t, err)
	assert.Len(t, quizzes, 1)
	assert.Equal(t, "Test Quiz", quizzes[0].Title)
}

func TestGetQuizByID(t *testing.T) {
	_, mock, service := setupTestDB(t)

	// Mock quiz query
	quizRows := sqlmock.NewRows([]string{"id", "title", "description", "deleted_at"}).
		AddRow(1, "Test Quiz", "Test Description", nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quizzes` WHERE `quizzes`.`id` = ? AND `quizzes`.`deleted_at` IS NULL ORDER BY `quizzes`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(quizRows)

	// Mock questions query
	questionRows := sqlmock.NewRows([]string{"id", "quiz_id", "text", "correct_answer", "deleted_at"}).
		AddRow(1, 1, "Test Question", "A", nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `questions` WHERE `questions`.`quiz_id` = ? AND `questions`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(questionRows)

	// Mock answers query
	answerRows := sqlmock.NewRows([]string{"id", "question_id", "text", "is_correct", "deleted_at"}).
		AddRow(1, 1, "Answer A", true, nil)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `answers` WHERE `answers`.`question_id` = ? AND `answers`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(answerRows)

	quiz, err := service.GetQuizByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, quiz)
	assert.Equal(t, "Test Quiz", quiz.Title)

	// Test case 2: Quiz not found
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `quizzes` WHERE `quizzes`.`id` = ? AND `quizzes`.`deleted_at` IS NULL ORDER BY `quizzes`.`id` LIMIT ?")).
		WithArgs(2, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	quiz, err = service.GetQuizByID(2)
	assert.Error(t, err)
	assert.Nil(t, quiz)
}

func TestCreateQuiz(t *testing.T) {
	_, mock, service := setupTestDB(t)

	quiz := &models.Quiz{
		Title:       "New Quiz",
		Description: "New Description",
		Category:    models.CategoryMath,
		Difficulty:  models.DifficultyEasy,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `quizzes`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.CreateQuiz(quiz)
	assert.NoError(t, err)
}

func TestGetQuizzesByCategory(t *testing.T) {
	_, mock, service := setupTestDB(t)

	rows := sqlmock.NewRows([]string{"id", "title", "category"}).
		AddRow(1, "Math Quiz", "math")

	mock.ExpectQuery("^SELECT (.+) FROM `quizzes`").
		WithArgs("math", true).
		WillReturnRows(rows)

	quizzes, err := service.GetQuizzesByCategory(models.CategoryMath)
	assert.NoError(t, err)
	assert.Len(t, quizzes, 1)
	assert.Equal(t, "Math Quiz", quizzes[0].Title)
}

func TestSubmitQuizResult(t *testing.T) {
	_, mock, service := setupTestDB(t)

	result := &models.Result{
		QuizID:         1,
		UserID:         1,
		Score:          85,
		TotalQuestions: 10,
		TimeTaken:      300,
	}

	// Mock quiz retrieval
	quizRows := sqlmock.NewRows([]string{"id", "title", "category"}).
		AddRow(1, "Test Quiz", "math")

	mock.ExpectQuery("^SELECT (.+) FROM `quizzes`").
		WithArgs(1).
		WillReturnRows(quizRows)

	// Mock transaction
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `results`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("^SELECT (.+) FROM `user_progress`").
		WithArgs(1, 1).
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectExec("^INSERT INTO `user_progress`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.SubmitQuizResult(result)
	assert.NoError(t, err)
}

func TestGetUserResults(t *testing.T) {
	_, mock, service := setupTestDB(t)

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "user_id", "score"}).
		AddRow(1, 1, 1, 85)

	mock.ExpectQuery("^SELECT (.+) FROM `results`").
		WithArgs(1).
		WillReturnRows(rows)

	results, err := service.GetUserResults(1)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, float64(85), results[0].Score)
}

func TestGetResultByID(t *testing.T) {
	_, mock, service := setupTestDB(t)

	rows := sqlmock.NewRows([]string{"id", "quiz_id", "user_id", "score"}).
		AddRow(1, 1, 1, 85)

	mock.ExpectQuery("^SELECT (.+) FROM `results`").
		WithArgs(1).
		WillReturnRows(rows)

	result, err := service.GetResultByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, float64(85), result.Score)
}
