package models

import (
	"fmt"

	"github.com/goravel/framework/database/orm"
)

type Player struct {
	orm.Model
	UserID     uint   `gorm:"column:user_id;not null;uniqueIndex" json:"user_id"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	DocumentID string `gorm:"column:document_id" json:"document_id"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Position   string `gorm:"column:position" json:"position"`
	Photo      string `gorm:"column:photo" json:"photo"`

	User     *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PhotoURL string `gorm:"-" json:"photo_url"`
}

func (p *Player) SetPhotoURL() {
	if p.Photo != "" && p.UpdatedAt != nil {
		p.PhotoURL = fmt.Sprintf("/players/%d/photo?t=%d", p.ID, p.UpdatedAt.Timestamp())
	}
}

func (*Player) TableName() string {
	return "players"
}
