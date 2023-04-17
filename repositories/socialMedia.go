package repositories

import (
	"github.com/alvinmdj/mygram-api/models"
	"gorm.io/gorm"
)

type SocialMediaRepoInterface interface {
	FindAll() (socialMedias []models.SocialMedia, err error)
	FindById(id int) (socialMedia models.SocialMedia, err error)
	Save(socialMedia models.SocialMedia) (models.SocialMedia, error)
	Update(socialMedia models.SocialMedia) (models.SocialMedia, error)
	Delete(socialMedia models.SocialMedia) (err error)
}

type SocialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepoInterface {
	return &SocialMediaRepo{
		db: db,
	}
}

func (s *SocialMediaRepo) FindAll() (socialMedias []models.SocialMedia, err error) {
	err = s.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("username", "id", "email", "age", "created_at", "updated_at")
	}).Find(&socialMedias).Error
	return
}

func (s *SocialMediaRepo) FindById(id int) (socialMedia models.SocialMedia, err error) {
	err = s.db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("username", "id", "email", "age", "created_at", "updated_at")
	}).First(&socialMedia, id).Error
	return
}

func (s *SocialMediaRepo) Save(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := s.db.Debug().Create(&socialMedia).Error
	return socialMedia, err
}

func (s *SocialMediaRepo) Update(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := s.db.Debug().Model(&socialMedia).
		Where("id = ?", socialMedia.ID).
		Updates(models.SocialMedia{
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
		}).Error
	return socialMedia, err
}

func (s *SocialMediaRepo) Delete(socialMedia models.SocialMedia) (err error) {
	err = s.db.Debug().Delete(&socialMedia).Error
	return
}
