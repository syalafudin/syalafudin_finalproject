package repositories

import (
	"github.com/alvinmdj/mygram-api/models"
	"gorm.io/gorm"
)

type PhotoRepoInterface interface {
	FindAll() (photos []models.Photo, err error)
	FindById(id int) (photo models.Photo, err error)
	Save(photo models.Photo) (models.Photo, error)
	Update(photo models.Photo) (models.Photo, error)
	Delete(photo models.Photo) (err error)
}

type PhotoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepoInterface {
	return &PhotoRepo{
		db: db,
	}
}

func (p *PhotoRepo) FindAll() (photos []models.Photo, err error) {
	err = p.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("username", "id", "email", "age", "created_at", "updated_at")
	}).Find(&photos).Error
	return
}

func (p *PhotoRepo) FindById(id int) (photo models.Photo, err error) {
	err = p.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("username", "id", "email", "age", "created_at", "updated_at")
	}).First(&photo, id).Error
	return
}

func (p *PhotoRepo) Save(photo models.Photo) (models.Photo, error) {
	err := p.db.Debug().Create(&photo).Error
	return photo, err
}

func (p *PhotoRepo) Update(photo models.Photo) (models.Photo, error) {
	err := p.db.Debug().Model(&photo).
		Where("id = ?", photo.ID).
		Updates(models.Photo{
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}).Error
	return photo, err
}

func (p *PhotoRepo) Delete(photo models.Photo) (err error) {
	err = p.db.Debug().Delete(&photo).Error
	return
}
