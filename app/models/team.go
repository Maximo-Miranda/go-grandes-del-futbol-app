package models

import "github.com/goravel/framework/database/orm"

type Team struct {
	orm.Model
	Name         string `gorm:"column:name;not null" json:"name"`
	Logo         string `gorm:"column:logo" json:"logo"`
	Color        string `gorm:"column:color" json:"color"`
	ContactPhone string `gorm:"column:contact_phone" json:"contact_phone"`

	Players []Player `gorm:"many2many:team_players" json:"players,omitempty"`
}

func (*Team) TableName() string {
	return "teams"
}
