package models

import "github.com/goravel/framework/database/orm"

type Venue struct {
	orm.Model
	Name        string `gorm:"column:name;not null" json:"name"`
	Address     string `gorm:"column:address" json:"address"`
	SurfaceType string `gorm:"column:surface_type" json:"surface_type"`
	Capacity    int    `gorm:"column:capacity;default:0" json:"capacity"`
}

func (*Venue) TableName() string {
	return "venues"
}
