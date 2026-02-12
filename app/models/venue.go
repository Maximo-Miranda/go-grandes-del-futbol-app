package models

import (
	"fmt"

	"github.com/goravel/framework/database/orm"
)

type Venue struct {
	orm.Model
	UUID    string  `gorm:"column:uuid;type:uuid;not null;uniqueIndex;default:(-)" json:"uuid"`
	Name    string  `gorm:"column:name;not null" json:"name"`
	Address string  `gorm:"column:address" json:"address"`
	Photo   string  `gorm:"column:photo" json:"photo"`
	Pitches []Pitch `gorm:"foreignKey:VenueID" json:"pitches,omitempty"`

	PhotoURL string `gorm:"-" json:"photo_url"`
}

func (v *Venue) SetPhotoURL() {
	if v.Photo != "" && v.UpdatedAt != nil {
		v.PhotoURL = fmt.Sprintf("/venues/%s/photo?t=%d", v.UUID, v.UpdatedAt.Timestamp())
	}
}

func (*Venue) TableName() string {
	return "venues"
}
