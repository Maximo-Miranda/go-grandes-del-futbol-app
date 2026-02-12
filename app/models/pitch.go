package models

import "github.com/goravel/framework/database/orm"

type Pitch struct {
	orm.Model
	VenueID     uint   `gorm:"column:venue_id;not null" json:"venue_id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	SurfaceType string `gorm:"column:surface_type;not null" json:"surface_type"`
	PitchType   string `gorm:"column:pitch_type;not null" json:"pitch_type"`
	Venue       *Venue `gorm:"foreignKey:VenueID" json:"venue,omitempty"`
}

func (*Pitch) TableName() string {
	return "pitches"
}
