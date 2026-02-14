package models

import (
	"fmt"

	"github.com/goravel/framework/database/orm"
)

type Team struct {
	orm.Model
	Name         string `gorm:"column:name;not null" json:"name"`
	Logo         string `gorm:"column:logo" json:"logo"`
	Color        string `gorm:"column:color;not null" json:"color"`
	ContactPhone string `gorm:"column:contact_phone;not null" json:"contact_phone"`

	Players []Player `gorm:"many2many:team_players" json:"players,omitempty"`
	LogoURL string   `gorm:"-" json:"logo_url"`
}

func (t *Team) SetLogoURL() {
	if t.Logo != "" && t.UpdatedAt != nil {
		t.LogoURL = fmt.Sprintf("/teams/%d/logo?t=%d", t.ID, t.UpdatedAt.Timestamp())
	}
}

func (*Team) TableName() string {
	return "teams"
}
