package enums

// QuizCategory represents the subject area of a quiz
type QuizCategory string

const (
	CategoryGeneral    QuizCategory = "general"
	CategoryScience    QuizCategory = "science"
	CategoryHistory    QuizCategory = "history"
	CategoryTechnology QuizCategory = "technology"
	CategoryMath       QuizCategory = "math"
	CategoryLanguage   QuizCategory = "language"
)

// ValidCategories returns all valid quiz categories
func ValidCategories() []QuizCategory {
	return []QuizCategory{
		CategoryGeneral,
		CategoryScience,
		CategoryHistory,
		CategoryTechnology,
		CategoryMath,
		CategoryLanguage,
	}
}

// IsValid checks if a category is valid
func (c QuizCategory) IsValid() bool {
	for _, validCategory := range ValidCategories() {
		if c == validCategory {
			return true
		}
	}
	return false
}

// String returns the string representation of the category
func (c QuizCategory) String() string {
	return string(c)
}
