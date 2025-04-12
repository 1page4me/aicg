package ranking

// LeaderboardEntry represents a user's position in the leaderboard
type LeaderboardEntry struct {
	UserID           uint   `json:"user_id"`
	Username         string `json:"username"`
	Score            int    `json:"score"`
	Rank             int    `json:"rank"`
	Category         string `json:"category"`
	AchievementCount int    `json:"achievement_count"`
}
