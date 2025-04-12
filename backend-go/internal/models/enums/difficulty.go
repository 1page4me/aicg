package enums

// QuizDifficulty represents the difficulty level of a quiz
type QuizDifficulty string

const (
	DifficultyEasy   QuizDifficulty = "easy"
	DifficultyMedium QuizDifficulty = "medium"
	DifficultyHard   QuizDifficulty = "hard"
)

// ValidDifficulties returns all valid quiz difficulties
func ValidDifficulties() []QuizDifficulty {
	return []QuizDifficulty{
		DifficultyEasy,
		DifficultyMedium,
		DifficultyHard,
	}
}

// IsValid checks if a difficulty level is valid
func (d QuizDifficulty) IsValid() bool {
	for _, validDifficulty := range ValidDifficulties() {
		if d == validDifficulty {
			return true
		}
	}
	return false
}

// String returns the string representation of the difficulty
func (d QuizDifficulty) String() string {
	return string(d)
}
