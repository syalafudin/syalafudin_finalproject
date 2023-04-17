package repositories

import (
	"github.com/alvinmdj/mygram-api/models"
	"gorm.io/gorm"
)

type CommentRepoInterface interface {
	FindAll(photoId int) (comments []models.Comment, err error)
	FindById(photoId int, commentId int) (comment models.Comment, err error)
	Save(comment models.Comment) (models.Comment, error)
	Update(comment models.Comment) (models.Comment, error)
	Delete(comment models.Comment) (err error)
}

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepoInterface {
	return &CommentRepo{
		db: db,
	}
}

func (co *CommentRepo) FindAll(photoId int) (comments []models.Comment, err error) {
	err = co.db.Debug().
		Where("photo_id = ?", photoId).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email", "age", "created_at", "updated_at")
		}).
		Find(&comments).Error
	return
}

func (co *CommentRepo) FindById(photoId int, commentId int) (comment models.Comment, err error) {
	err = co.db.Debug().
		Where("photo_id = ?", photoId).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email", "age", "created_at", "updated_at")
		}).
		First(&comment, commentId).Error
	return
}

func (co *CommentRepo) Save(comment models.Comment) (models.Comment, error) {
	err := co.db.Debug().Create(&comment).Error
	return comment, err
}

func (co *CommentRepo) Update(comment models.Comment) (models.Comment, error) {
	err := co.db.Debug().Model(&comment).
		Where("id = ?", comment.ID).
		Updates(models.Comment{
			Message: comment.Message,
		}).Error
	return comment, err
}

func (co *CommentRepo) Delete(comment models.Comment) (err error) {
	err = co.db.Debug().Delete(&comment).Error
	return
}
