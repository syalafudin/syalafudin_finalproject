package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Base
	Name           string `gorm:"not null"`
	SocialMediaURL string `gorm:"not null;uniqueIndex"`
	UserID         uint
	User           User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	input := SocialMediaCreateInput{
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
		UserID:         s.UserID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	// validate input
	input := SocialMediaUpdateInput{
		ID:             s.ID,
		Name:           s.Name,
		SocialMediaURL: s.SocialMediaURL,
		UserID:         s.UserID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}
