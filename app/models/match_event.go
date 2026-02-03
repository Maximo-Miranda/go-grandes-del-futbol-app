package models

import "github.com/goravel/framework/database/orm"

type MatchEvent struct {
	orm.Model
	MatchID   uint   `gorm:"column:match_id;not null" json:"match_id"`
	PlayerID  uint   `gorm:"column:player_id;not null" json:"player_id"`
	TeamID    uint   `gorm:"column:team_id;not null" json:"team_id"`
	EventType string `gorm:"column:event_type;not null" json:"event_type"`
	Minute    int    `gorm:"column:minute;default:0" json:"minute"`

	Match  *Match  `gorm:"foreignKey:MatchID" json:"match,omitempty"`
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
	Team   *Team   `gorm:"foreignKey:TeamID" json:"team,omitempty"`
}

func (*MatchEvent) TableName() string {
	return "match_events"
}
