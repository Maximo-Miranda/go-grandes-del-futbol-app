package models

import "github.com/goravel/framework/database/orm"

type TeamPlayer struct {
	orm.Model
	TeamID       uint `gorm:"column:team_id;not null" json:"team_id"`
	PlayerID     uint `gorm:"column:player_id;not null" json:"player_id"`
	JerseyNumber int  `gorm:"column:jersey_number;default:0" json:"jersey_number"`
	IsCaptain    bool `gorm:"column:is_captain;default:false" json:"is_captain"`

	Team   *Team   `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

func (*TeamPlayer) TableName() string {
	return "team_players"
}
