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

type CommentHdlInterface interface {
	GetAll(c *gin.Context)
	GetOneById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CommentHandler struct {
	commentSvc services.CommentSvcInterface
}

func NewCommentHdl(commentSvc services.CommentSvcInterface) CommentHdlInterface {
	return &CommentHandler{
		commentSvc: commentSvc,
	}
}

// Comments GetAll godoc
// @Summary Get all comments associated with the photo id
// @Description Get all comments associated with the photo id
// @Tags comments
// @Param photoId path string true "get comment associated with the photo id"
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} []models.CommentGetOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId}/comments [get]
func (co *CommentHandler) GetAll(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	comments, err := co.commentSvc.GetAll(photoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	commentsResponse := []models.CommentGetOutput{}
	for _, comment := range comments {
		commentOutput := models.CommentGetOutput{
			Base:    comment.Base,
			Message: comment.Message,
			User: models.UserRegisterOutput{
				Base:     comment.User.Base,
				Username: comment.User.Username,
				Email:    comment.User.Email,
				Age:      comment.User.Age,
			},
		}
		commentsResponse = append(commentsResponse, commentOutput)
	}
	c.JSON(http.StatusOK, commentsResponse)
}

// Comment GetOneById godoc
// @Summary Get one comment by id
// @Description Get one comment by id
// @Tags comments
// @Param photoId path string true "get comment associated with the photo id"
// @Param commentId path string true "get comment by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Produce json
// @Success 200 {object} models.CommentGetOutput{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId}/comments/{commentId} [get]
func (co *CommentHandler) GetOneById(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	comment, err := co.commentSvc.GetOneById(photoId, commentId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	commentResponse := models.CommentGetOutput{
		Base:    comment.Base,
		Message: comment.Message,
		User: models.UserRegisterOutput{
			Base:     comment.User.Base,
			Username: comment.User.Username,
			Email:    comment.User.Email,
			Age:      comment.User.Age,
		},
	}
	c.JSON(http.StatusOK, commentResponse)
}

// Comment Create godoc
// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Accept json,mpfd
// @Produce json
// @Param photoId path string true "create comment associated with the photo id"
// @Param models.CommentCreateInput body models.CommentCreateInputSwagger{} true "create comment"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 201 {object} models.CommentCreateOutput{}
// @Failure 400 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId}/comments [post]
func (co *CommentHandler) Create(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	commentInput := models.CommentCreateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// get photo id from path param
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	commentInput.UserID = userId
	commentInput.PhotoID = uint(photoId)

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&commentInput)
	} else {
		c.ShouldBind(&commentInput)
	}

	comment, err := co.commentSvc.Create(commentInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "BAD REQUEST",
			Message: err.Error(),
		})
		return
	}

	commentResponse := models.CommentCreateOutput{
		Base:    comment.Base,
		Message: comment.Message,
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
	}
	c.JSON(http.StatusCreated, commentResponse)
}

// Comment Update godoc
// @Summary Update comment
// @Description Update comment
// @Tags comments
// @Accept json,mpfd
// @Produce json
// @Param photoId path string true "update comment associated with the photo id"
// @Param commentId path string true "update comment by id"
// @Param models.CommentUpdateInput body models.CommentUpdateInputSwagger{} true "update comment"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.CommentUpdateOutput{}
// @Failure 403 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId}/comments/{commentId} [put]
func (co *CommentHandler) Update(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	contentType := helpers.GetContentType(c)
	commentInput := models.CommentUpdateInput{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// store id and user id to input struct
	commentInput.ID = uint(commentId)
	commentInput.UserID = userId
	commentInput.PhotoID = uint(photoId)

	if contentType == helpers.AppJson {
		c.ShouldBindJSON(&commentInput)
	} else {
		c.ShouldBind(&commentInput)
	}

	comment, err := co.commentSvc.Update(commentInput)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	commentResponse := models.CommentUpdateOutput{
		Base:    comment.Base,
		Message: comment.Message,
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
	}
	c.JSON(http.StatusOK, commentResponse)
}

// Comment Delete godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags comments
// @Produce json
// @Param photoId path string true "delete comment associated with the photo id"
// @Param commentId path string true "delete comment by id"
// @Param Authorization header string true "format: Bearer token-here"
// @Success 200 {object} models.DeleteResponse{}
// @Failure 403 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Router /api/v1/photos/{photoId}/comments/{commentId} [delete]
func (co *CommentHandler) Delete(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if err := co.commentSvc.Delete(commentId); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT FOUND",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DeleteResponse{
		Message: fmt.Sprintf("comment data with id %d has been deleted", commentId),
	})
}
