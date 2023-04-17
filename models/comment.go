package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Base
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	UserID  uint
	PhotoID uint
	User    User
	Photo   Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	input := CommentCreateInput{
		Message: c.Message,
		UserID:  c.UserID,
		PhotoID: c.PhotoID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	// validate input
	input := CommentUpdateInput{
		ID:      c.ID,
		Message: c.Message,
		UserID:  c.UserID,
		PhotoID: c.PhotoID,
	}
	_, err = govalidator.ValidateStruct(input)
	return
}
