package models

import "github.com/goravel/framework/database/orm"

type TournamentTeam struct {
	orm.Model
	TournamentID uint  `gorm:"column:tournament_id;not null" json:"tournament_id"`
	TeamID       uint  `gorm:"column:team_id;not null" json:"team_id"`
	GroupID      *uint `gorm:"column:group_id" json:"group_id"`

	Tournament *Tournament `gorm:"foreignKey:TournamentID" json:"tournament,omitempty"`
	Team       *Team       `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Group      *Group      `gorm:"foreignKey:GroupID" json:"group,omitempty"`
}

func (*TournamentTeam) TableName() string {
	return "tournament_teams"
}
