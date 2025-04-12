package enums

// QuestionType defines the type of question
type QuestionType string

const (
	QuestionTypeMultipleChoice QuestionType = "multiple_choice"
	QuestionTypeTrueFalse      QuestionType = "true_false"
	QuestionTypeShortAnswer    QuestionType = "short_answer"
	QuestionTypeEssay          QuestionType = "essay"
)

// ValidQuestionTypes returns all valid question types
func ValidQuestionTypes() []QuestionType {
	return []QuestionType{
		QuestionTypeMultipleChoice,
		QuestionTypeTrueFalse,
		QuestionTypeShortAnswer,
		QuestionTypeEssay,
	}
}

// IsValid checks if a question type is valid
func (t QuestionType) IsValid() bool {
	for _, validType := range ValidQuestionTypes() {
		if t == validType {
			return true
		}
	}
	return false
}

// String returns the string representation of the question type
func (t QuestionType) String() string {
	return string(t)
}
