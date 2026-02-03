package models

import "github.com/goravel/framework/database/orm"

type Player struct {
	orm.Model
	UserID     *uint  `gorm:"column:user_id" json:"user_id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	DocumentID string `gorm:"column:document_id" json:"document_id"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Position   string `gorm:"column:position" json:"position"`
	Photo      string `gorm:"column:photo" json:"photo"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (*Player) TableName() string {
	return "players"
}
