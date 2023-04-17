package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Base
	Title    string `gorm:"not null"`
	Caption  string `gorm:"not null"`
	PhotoURL string `gorm:"not null"`
	UserID   uint
	User     User
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	input := PhotoCreateInput{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
		UserID:   p.UserID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	// validate input
	input := PhotoUpdateInput{
		ID:       p.ID,
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
		UserID:   p.UserID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}
