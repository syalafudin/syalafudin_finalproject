package handlers

import (
	"net/http"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
)

type UserHdlInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type UserHandler struct {
	userSvc services.UserSvcInterface
}

func NewUserHdl(userSvc services.UserSvcInterface) UserHdlInterface {
	return &UserHandler{
		userSvc: userSvc,
	}
}

// User Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags users
// @Accept json,mpfd
// @Produce json
// @Param models.UserRegisterInput body models.UserRegisterInput{} true "register user"
// @Success 201 {object} models.UserRegisterOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/users/register [post]
func (u *UserHandler) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userInput := models.UserRegisterInput{}

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&userInput)
	} else {
		c.ShouldBind(&userInput)
	}

	user, err := u.userSvc.Register(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	userResponse := models.UserRegisterOutput{
		Base:     user.Base,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
	c.JSON(http.StatusCreated, userResponse)
}

// User Login godoc
// @Summary User login
// @Description User login
// @Tags users
// @Accept json,mpfd
// @Produce json
// @Param models.UserLoginInput body models.UserLoginInput{} true "login user"
// @Success 201 {object} models.UserLoginOutput{}
// @Failure 401 {object} models.ErrorResponse{}
// @Router /api/v1/users/login [post]
func (u *UserHandler) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userInput := models.UserLoginInput{}

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&userInput)
	} else {
		c.ShouldBind(&userInput)
	}

	token, err := u.userSvc.Login(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error:   "UNAUTHORIZED",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.UserLoginOutput{
		Token: token,
	})
}
