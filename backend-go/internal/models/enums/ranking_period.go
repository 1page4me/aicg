package enums

// RankingPeriod represents the time period for rankings
type RankingPeriod string

const (
	RankingPeriodWeekly  RankingPeriod = "weekly"
	RankingPeriodMonthly RankingPeriod = "monthly"
	RankingPeriodAllTime RankingPeriod = "all_time"
)

// ValidRankingPeriods returns all valid ranking periods
func ValidRankingPeriods() []RankingPeriod {
	return []RankingPeriod{
		RankingPeriodWeekly,
		RankingPeriodMonthly,
		RankingPeriodAllTime,
	}
}

// IsValid checks if a ranking period is valid
func (p RankingPeriod) IsValid() bool {
	for _, validPeriod := range ValidRankingPeriods() {
		if p == validPeriod {
			return true
		}
	}
	return false
}

// String returns the string representation of the ranking period
func (p RankingPeriod) String() string {
	return string(p)
}
