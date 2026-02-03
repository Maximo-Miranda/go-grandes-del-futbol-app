package models

import "github.com/goravel/framework/database/orm"

type MatchLineup struct {
	orm.Model
	MatchID   uint `gorm:"column:match_id;not null" json:"match_id"`
	TeamID    uint `gorm:"column:team_id;not null" json:"team_id"`
	PlayerID  uint `gorm:"column:player_id;not null" json:"player_id"`
	IsStarter bool `gorm:"column:is_starter;default:true" json:"is_starter"`

	Match  *Match  `gorm:"foreignKey:MatchID" json:"match,omitempty"`
	Team   *Team   `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

func (*MatchLineup) TableName() string {
	return "match_lineups"
}
