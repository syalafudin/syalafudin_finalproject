package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SocialMediaHdlInterface interface {
	GetAll(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type SocialMediaHandler struct {
	socialMediaSvc services.SocialMediaSvcInterface
}

func NewSocialMediaHdl(socialMediaSvc services.SocialMediaSvcInterface) SocialMediaHdlInterface {
	return &SocialMediaHandler{
		socialMediaSvc: socialMediaSvc,
	}
}

// Social Media GetAll godoc
// @Summary Get all social media
// @Description Get all social media
// @Tags socialMedias
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} []models.SocialMediaGetOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/social-medias [get]
func (s *SocialMediaHandler) GetAll(c *gin.Context) {
	socialMedias, err := s.socialMediaSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	socialMediasResponse := []models.SocialMediaGetOutput{}
	for _, socialMedia := range socialMedias {
		socialMediaOutput := models.SocialMediaGetOutput{
			Base:           socialMedia.Base,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			User: models.UserRegisterOutput{
				Base:     socialMedia.User.Base,
				Username: socialMedia.User.Username,
				Email:    socialMedia.User.Email,
				Age:      socialMedia.User.Age,
			},
		}
		socialMediasResponse = append(socialMediasResponse, socialMediaOutput)
	}
	c.JSON(http.StatusOK, socialMediasResponse)
}

// Social Media GetOneById godoc
// @Summary Get one social media by id
// @Description Get one social media by id
// @Tags socialMedias
// @Param socialMediaId path string true "get social media by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} models.SocialMediaGetOutput{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/social-medias/{socialMediaId} [get]
func (s *SocialMediaHandler) GetOneById(c *gin.Context) {
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	socialMedia, err := s.socialMediaSvc.GetOneById(socialMediaId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	socialMediaResponse := models.SocialMediaGetOutput{
		Base:           socialMedia.Base,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		User: models.UserRegisterOutput{
			Base:     socialMedia.User.Base,
			Username: socialMedia.User.Username,
			Email:    socialMedia.User.Email,
			Age:      socialMedia.User.Age,
		},
	}
	c.JSON(http.StatusOK, socialMediaResponse)
}

// Social Media Create godoc
// @Summary Create social media
// @Description Create social media
// @Tags socialMedias
// @Accept json,mpfd
// @Produce json
// @Param models.SocialMediaCreateInput body models.SocialMediaCreateInputSwagger{} true "create social media"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 201 {object} models.SocialMediaCreateOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/social-medias [post]
func (s *SocialMediaHandler) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMediaInput := models.SocialMediaCreateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMediaInput.UserID = userId

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&socialMediaInput)
	} else {
		c.ShouldBind(&socialMediaInput)
	}

	socialMedia, err := s.socialMediaSvc.Create(socialMediaInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	socialMediaResponse := models.SocialMediaCreateOutput{
		Base:           socialMedia.Base,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
	}
	c.JSON(http.StatusCreated, socialMediaResponse)
}

// Social Media Update godoc
// @Summary Update social media
// @Description Update social media
// @Tags socialMedias
// @Accept json,mpfd
// @Produce json
// @Param socialMediaId path string true "update social media by id"
// @Param models.SocialMediaUpdateInput body models.SocialMediaUpdateInputSwagger{} true "update social media"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.SocialMediaUpdateOutput{}
// @Failure 403 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/social-medias/{socialMediaId} [put]
func (s *SocialMediaHandler) Update(c *gin.Context) {
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	contentType := helpers.GetContentType(c)
	socialMediaInput := models.SocialMediaUpdateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// store id and user id to input struct
	socialMediaInput.ID = uint(socialMediaId)
	socialMediaInput.UserID = userId

	// get req body: name and social media url
	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&socialMediaInput)
	} else {
		c.ShouldBind(&socialMediaInput)
	}

	socialMedia, err := s.socialMediaSvc.Update(socialMediaInput)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	socialMediaResponse := models.SocialMediaUpdateOutput{
		Base:           socialMedia.Base,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
	}
	c.JSON(http.StatusOK, socialMediaResponse)
}

// Social Media Delete godoc
// @Summary Delete social media
// @Description Delete social media
// @Tags socialMedias
// @Produce json
// @Param socialMediaId path string true "delete social media by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.DeleteResponse{}
// @Failure 403 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/social-medias/{socialMediaId} [delete]
func (s *SocialMediaHandler) Delete(c *gin.Context) {
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	if err := s.socialMediaSvc.Delete(socialMediaId); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DeleteResponse{
		Message: fmt.Sprintf("social media data with id %d has been deleted", socialMediaId),
	})
}
