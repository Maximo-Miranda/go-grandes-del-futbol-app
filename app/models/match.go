package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Match struct {
	orm.Model
	TournamentID uint       `gorm:"column:tournament_id;not null" json:"tournament_id"`
	GroupID      *uint      `gorm:"column:group_id" json:"group_id"`
	HomeTeamID   uint       `gorm:"column:home_team_id;not null" json:"home_team_id"`
	AwayTeamID   uint       `gorm:"column:away_team_id;not null" json:"away_team_id"`
	VenueID      *uint      `gorm:"column:venue_id" json:"venue_id"`
	MatchDate    *time.Time `gorm:"column:match_date" json:"match_date"`
	Status       string     `gorm:"column:status;default:scheduled;not null" json:"status"`
	HomeScore    int        `gorm:"column:home_score;default:0" json:"home_score"`
	AwayScore    int        `gorm:"column:away_score;default:0" json:"away_score"`
	Round        int        `gorm:"column:round;default:1" json:"round"`

	Tournament *Tournament `gorm:"foreignKey:TournamentID" json:"tournament,omitempty"`
	Group      *Group      `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	HomeTeam   *Team       `gorm:"foreignKey:HomeTeamID" json:"home_team,omitempty"`
	AwayTeam   *Team       `gorm:"foreignKey:AwayTeamID" json:"away_team,omitempty"`
	Venue      *Venue      `gorm:"foreignKey:VenueID" json:"venue,omitempty"`
	Events     []MatchEvent  `gorm:"foreignKey:MatchID" json:"events,omitempty"`
	Lineups    []MatchLineup `gorm:"foreignKey:MatchID" json:"lineups,omitempty"`
}

func (*Match) TableName() string {
	return "matches"
}
