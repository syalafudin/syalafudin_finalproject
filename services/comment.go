package services

import (
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
)

type CommentSvcInterface interface {
	GetAll(photoId int) (comments []models.Comment, err error)
	GetOneById(photoId int, commentId int) (comment models.Comment, err error)
	Create(commentInput models.CommentCreateInput) (comment models.Comment, err error)
	Update(commentInput models.CommentUpdateInput) (comment models.Comment, err error)
	Delete(commentId int) (err error)
}

type CommentSvc struct {
	commentRepo repositories.CommentRepoInterface
}

func NewCommentSvc(commentRepo repositories.CommentRepoInterface) CommentSvcInterface {
	return &CommentSvc{
		commentRepo: commentRepo,
	}
}

func (co *CommentSvc) GetAll(photoId int) (comments []models.Comment, err error) {
	comments, err = co.commentRepo.FindAll(photoId)
	return
}

func (co *CommentSvc) GetOneById(photoId int, commentId int) (comment models.Comment, err error) {
	comment, err = co.commentRepo.FindById(photoId, commentId)
	return
}

func (co *CommentSvc) Create(commentInput models.CommentCreateInput) (comment models.Comment, err error) {
	comment = models.Comment{
		Message: commentInput.Message,
		UserID:  commentInput.UserID,
		PhotoID: commentInput.PhotoID,
	}

	comment, err = co.commentRepo.Save(comment)
	return
}

func (co *CommentSvc) Update(commentInput models.CommentUpdateInput) (comment models.Comment, err error) {
	comment = models.Comment{
		Base:    models.Base{ID: commentInput.ID},
		Message: commentInput.Message,
		UserID:  commentInput.UserID,
		PhotoID: commentInput.PhotoID,
	}

	comment, err = co.commentRepo.Update(comment)
	return
}

func (co *CommentSvc) Delete(commentId int) (err error) {
	comment := models.Comment{
		Base: models.Base{ID: uint(commentId)},
	}

	err = co.commentRepo.Delete(comment)
	return
}
