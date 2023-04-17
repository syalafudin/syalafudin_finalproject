package services

import (
	"errors"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
)

type UserSvcInterface interface {
	Register(userInput models.UserRegisterInput) (user models.User, err error)
	Login(userInput models.UserLoginInput) (token string, err error)
}

type UserSvc struct {
	userRepo repositories.UserRepoInterface
}

func NewUserSvc(userRepo repositories.UserRepoInterface) UserSvcInterface {
	return &UserSvc{
		userRepo: userRepo,
	}
}

func (u *UserSvc) Register(userInput models.UserRegisterInput) (user models.User, err error) {
	user = models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
		Age:      userInput.Age,
	}

	user, err = u.userRepo.Save(user)
	return
}

func (u *UserSvc) Login(userInput models.UserLoginInput) (token string, err error) {
	user := models.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	user, err = u.userRepo.FindByEmail(user)
	if err != nil {
		err = errors.New("invalid email or password")
		return
	}

	if isEqual := helpers.CompareHash([]byte(user.Password), []byte(userInput.Password)); !isEqual {
		err = errors.New("invalid email or password")
		return
	}

	token = helpers.GenerateToken(user.ID, user.Email)
	return
}
