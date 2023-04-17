package models

type UserRegisterInput struct {
	Username string `json:"username" form:"username" valid:"required~username is required"`
	Email    string `json:"email" form:"email" valid:"required~email is required,email~invalid email format"`
	Password string `json:"password" form:"password" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
	Age      int    `json:"age" form:"age" valid:"required~age is required,range(8|99)~user must be at least 8 years old"`
}

type UserRegisterOutput struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserLoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginOutput struct {
	Token string `json:"token"`
}
