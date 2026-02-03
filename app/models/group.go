package models

import "github.com/goravel/framework/database/orm"

type Group struct {
	orm.Model
	TournamentID uint   `gorm:"column:tournament_id;not null" json:"tournament_id"`
	Name         string `gorm:"column:name;not null" json:"name"`

	Tournament *Tournament `gorm:"foreignKey:TournamentID" json:"tournament,omitempty"`
}

func (*Group) TableName() string {
	return "groups"
}
