package models

type CommentGetOutput struct {
	Base
	Message string             `json:"message"`
	User    UserRegisterOutput `json:"user"`
}

type CommentCreateInput struct {
	Message string `json:"message" form:"message" valid:"required~message is required"`
	UserID  uint   `valid:"required~user ID is required"`
	PhotoID uint   `valid:"required~photo ID is required"`
}

type CommentCreateInputSwagger struct {
	Message string `json:"message" form:"message"`
}

type CommentCreateOutput struct {
	Base
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

type CommentUpdateInput struct {
	ID      uint   `valid:"required~ID is required"`
	Message string `json:"message" form:"message" valid:"required~message is required"`
	UserID  uint   `valid:"required~user ID is required"`
	PhotoID uint   `valid:"required~photo ID is required"`
}

type CommentUpdateInputSwagger = CommentCreateInputSwagger

type CommentUpdateOutput = CommentCreateOutput
