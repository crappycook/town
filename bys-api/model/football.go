package model

import "time"

type MatchResult string

const (
	MatchResultInProgress = MatchResult("in-progress")
	MatchResultHome       = MatchResult("home")
	MatchResultAway       = MatchResult("away")
	MatchResultDraw       = MatchResult("draw")
)

type FootballMatch struct {
	TableBase

	HomeTeam  string      `gorm:"not null;type:char(8)" json:"home_team"`
	AwayTeam  string      `gorm:"not null;type:char(8)" json:"away_team"`
	HomeScore uint8       `gorm:"not null;type:tinyint" json:"home_score"`
	AwayScore uint8       `gorm:"not null;type:tinyint" json:"away_score"`
	Result    MatchResult `gorm:"not null;type:char(16)" json:"result"`
	StartTime time.Time   `gorm:"not null;type:datetime(3)" json:"start_time"`
}
