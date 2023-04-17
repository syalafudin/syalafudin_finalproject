package models

import (
	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username     string        `gorm:"not null;uniqueIndex"`
	Email        string        `gorm:"not null;uniqueIndex"`
	Password     string        `gorm:"not null"`
	Age          int           `gorm:"not null"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	input := UserRegisterInput{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
	}
	if _, err = govalidator.ValidateStruct(input); err != nil {
		return
	}

	// hash password
	u.Password, err = helpers.HashPassword(u.Password)
	return
}
