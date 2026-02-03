package models

import "github.com/goravel/framework/database/orm"

type Standing struct {
	orm.Model
	TournamentID   uint  `gorm:"column:tournament_id;not null" json:"tournament_id"`
	GroupID        *uint `gorm:"column:group_id" json:"group_id"`
	TeamID         uint  `gorm:"column:team_id;not null" json:"team_id"`
	Played         int   `gorm:"column:played;default:0" json:"played"`
	Won            int   `gorm:"column:won;default:0" json:"won"`
	Drawn          int   `gorm:"column:drawn;default:0" json:"drawn"`
	Lost           int   `gorm:"column:lost;default:0" json:"lost"`
	GoalsFor       int   `gorm:"column:goals_for;default:0" json:"goals_for"`
	GoalsAgainst   int   `gorm:"column:goals_against;default:0" json:"goals_against"`
	GoalDifference int   `gorm:"column:goal_difference;default:0" json:"goal_difference"`
	Points         int   `gorm:"column:points;default:0" json:"points"`

	Tournament *Tournament `gorm:"foreignKey:TournamentID" json:"tournament,omitempty"`
	Group      *Group      `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	Team       *Team       `gorm:"foreignKey:TeamID" json:"team,omitempty"`
}

func (*Standing) TableName() string {
	return "standings"
}
