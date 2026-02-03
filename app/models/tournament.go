package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Tournament struct {
	orm.Model
	Name        string     `gorm:"column:name;not null" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Format      string     `gorm:"column:format;default:round_robin;not null" json:"format"`
	GameType    string     `gorm:"column:game_type;default:5v5;not null" json:"game_type"`
	Status      string     `gorm:"column:status;default:draft;not null" json:"status"`
	StartDate   *time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate     *time.Time `gorm:"column:end_date" json:"end_date"`
	VenueID     *uint      `gorm:"column:venue_id" json:"venue_id"`

	Venue  *Venue  `gorm:"foreignKey:VenueID" json:"venue,omitempty"`
	Groups []Group `gorm:"foreignKey:TournamentID" json:"groups,omitempty"`
	Teams  []Team  `gorm:"many2many:tournament_teams" json:"teams,omitempty"`
}

func (*Tournament) TableName() string {
	return "tournaments"
}
