package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PhotoHdlInterface interface {
	GetAll(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PhotoHandler struct {
	photoSvc services.PhotoSvcInterface
}

func NewPhotoHdl(photoSvc services.PhotoSvcInterface) PhotoHdlInterface {
	return &PhotoHandler{
		photoSvc: photoSvc,
	}
}

// Photo GetAll godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags photos
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} []models.PhotoGetOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/photos [get]
func (p *PhotoHandler) GetAll(c *gin.Context) {
	photos, err := p.photoSvc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	photosResponse := []models.PhotoGetOutput{}
	for _, photo := range photos {
		photoOutput := models.PhotoGetOutput{
			Base:     photo.Base,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			User: models.UserRegisterOutput{
				Base:     photo.User.Base,
				Username: photo.User.Username,
				Email:    photo.User.Email,
				Age:      photo.User.Age,
			},
		}
		photosResponse = append(photosResponse, photoOutput)
	}
	c.JSON(http.StatusOK, photosResponse)
}

// Photo GetOneById godoc
// @Summary Get one photo by id
// @Description Get one photo by id
// @Tags photos
// @Param photoId path string true "get photo by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} models.PhotoGetOutput{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId} [get]
func (p *PhotoHandler) GetOneById(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	photo, err := p.photoSvc.GetOneById(photoId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	photoResponse := models.PhotoGetOutput{
		Base:     photo.Base,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		User: models.UserRegisterOutput{
			Base:     photo.User.Base,
			Username: photo.User.Username,
			Email:    photo.User.Email,
			Age:      photo.User.Age,
		},
	}
	c.JSON(http.StatusOK, photoResponse)
}

// Photo Create godoc
// @Summary Create photos
// @Description Create photos
// @Tags photos
// @Accept mpfd
// @Produce json
// @Param models.PhotoCreateInput formData models.PhotoCreateInputSwagger true "create photo"
// @Param photo formData file true "upload photo"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 201 {object} models.PhotoCreateOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Failure 413 {object} models.ErrorResponse{}
// @Router /api/v1/photos [post]
func (p *PhotoHandler) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	photoInput := models.PhotoCreateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	photoInput.UserID = userId

	// only accept multipart/form-data
	if contentType == helpers.AppJson {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: "invalid content type",
		})
		return
	} else {
		c.ShouldBind(&photoInput)
	}

	// photo source, check if photo is uploaded
	photoFileHeader, err := c.FormFile("photo")
	if err != nil {
		log.Printf("get form err - %s", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: "no photo file uploaded",
		})
		return
	}

	// Check if the file is an image
	ext := filepath.Ext(photoFileHeader.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: "invalid file type",
		})
		return
	}

	photo, err := p.photoSvc.Create(photoInput, photoFileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	photoResponse := models.PhotoCreateOutput{
		Base:     photo.Base,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}
	c.JSON(http.StatusCreated, photoResponse)
}

// Photo Update godoc
// @Summary Update photo
// @Description Update photo
// @Tags photos
// @Accept json,mpfd
// @Produce json
// @Param photoId path string true "update photo by id"
// @Param models.PhotoUpdateInput formData models.PhotoUpdateInputSwagger true "update photo"
// @Param photo formData file false "upload photo"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.PhotoUpdateOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Failure 403 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId} [put]
func (p *PhotoHandler) Update(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	contentType := helpers.GetContentType(c)
	photoInput := models.PhotoUpdateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// store id and user id to input struct
	photoInput.ID = uint(photoId)
	photoInput.UserID = userId

	// only accept multipart/form-data
	if contentType == helpers.AppJson {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: "invalid content type",
		})
		return
	} else {
		c.ShouldBind(&photoInput)
	}

	// photo source, check if photo is uploaded
	// new photo is not mandatory for update
	photoFileHeader, _ := c.FormFile("photo")
	if photoFileHeader != nil {
		// Check if the file is an image
		ext := filepath.Ext(photoFileHeader.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "BAD REQUEST",
				Message: "invalid file type",
			})
			return
		}
	}

	photo, err := p.photoSvc.Update(photoInput, photoFileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	photoResponse := models.PhotoUpdateOutput{
		Base:     photo.Base,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}
	c.JSON(http.StatusOK, photoResponse)
}

// Photo Delete godoc
// @Summary Delete photo
// @Description Delete photo
// @Tags photos
// @Produce json
// @Param photoId path string true "delete photo by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.DeleteResponse{}
// @Failure 403 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId} [delete]
func (p *PhotoHandler) Delete(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if err := p.photoSvc.Delete(photoId); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DeleteResponse{
		Message: fmt.Sprintf("photo data with id %d has been deleted", photoId),
	})
}
