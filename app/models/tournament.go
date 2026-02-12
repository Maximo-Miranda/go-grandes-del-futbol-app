package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Tournament struct {
	orm.Model
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Format      string    `gorm:"column:format;not null" json:"format"`
	GameType    string    `gorm:"column:game_type;not null" json:"game_type"`
	Status      string    `gorm:"column:status;not null" json:"status"`
	StartDate   time.Time `gorm:"column:start_date;not null" json:"start_date"`
	EndDate     time.Time `gorm:"column:end_date;not null" json:"end_date"`
	VenueID     uint      `gorm:"column:venue_id;not null" json:"venue_id"`

	Venue  *Venue  `gorm:"foreignKey:VenueID" json:"venue,omitempty"`
	Groups []Group `gorm:"foreignKey:TournamentID" json:"groups,omitempty"`
	Teams  []Team  `gorm:"many2many:tournament_teams" json:"teams,omitempty"`
}

func (*Tournament) TableName() string {
	return "tournaments"
}
