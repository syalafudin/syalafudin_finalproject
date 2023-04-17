package models

type PhotoGetOutput struct {
	Base
	Title    string             `json:"title"`
	Caption  string             `json:"caption"`
	PhotoURL string             `json:"photo_url"`
	User     UserRegisterOutput `json:"user"`
}

type PhotoCreateInput struct {
	Title    string `form:"title" valid:"required~title is required"`
	Caption  string `form:"caption" valid:"required~caption is required"`
	PhotoURL string `form:"photo_url" valid:"required~photo URL is required"`
	UserID   uint   `valid:"required~user ID is required"`
}

// this struct only used for swagger docs to generate desired input
type PhotoCreateInputSwagger struct {
	Title   string `form:"title"`
	Caption string `form:"caption"`
}

type PhotoCreateOutput struct {
	Base
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type PhotoUpdateInput struct {
	ID       uint   `valid:"required~ID is required"`
	Title    string `form:"title" valid:"required~title is required"`
	Caption  string `form:"caption" valid:"required~caption is required"`
	PhotoURL string `form:"photo_url" valid:"required~photo URL is required"`
	UserID   uint   `valid:"required~user ID is required"`
}

type PhotoUpdateInputSwagger = PhotoCreateInputSwagger

type PhotoUpdateOutput = PhotoCreateOutput
