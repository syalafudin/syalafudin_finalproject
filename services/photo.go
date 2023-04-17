package services

import (
	"log"
	"mime/multipart"

	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/models"
	"github.com/alvinmdj/mygram-api/repositories"
	"github.com/asaskevich/govalidator"
)

type PhotoSvcInterface interface {
	GetAll() (photos []models.Photo, err error)
	GetOneById(id int) (photo models.Photo, err error)
	Create(photoInput models.PhotoCreateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error)
	Update(photoInput models.PhotoUpdateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error)
	Delete(id int) (err error)
}

type PhotoSvc struct {
	photoRepo repositories.PhotoRepoInterface
}

func NewPhotoSvc(photoRepo repositories.PhotoRepoInterface) PhotoSvcInterface {
	return &PhotoSvc{
		photoRepo: photoRepo,
	}
}

func (p *PhotoSvc) GetAll() (photos []models.Photo, err error) {
	photos, err = p.photoRepo.FindAll()
	return
}

func (p *PhotoSvc) GetOneById(id int) (photo models.Photo, err error) {
	photo, err = p.photoRepo.FindById(id)
	return
}

func (p *PhotoSvc) Create(photoInput models.PhotoCreateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error) {
	// validate other input before upload file to cloudinary
	photoInput.PhotoURL = "placeholder"
	_, err = govalidator.ValidateStruct(photoInput)
	if err != nil {
		return
	}

	// open the file and get its content
	photoFile, err := photoFileHeader.Open()
	if err != nil {
		log.Printf("error opening file: %v", err)
		return
	}
	defer photoFile.Close()

	// upload file to cloudinary
	photoUrl, err := helpers.UploadToCloudinary(photoFile)
	if err != nil {
		return
	}

	photo = models.Photo{
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		UserID:   photoInput.UserID,
		PhotoURL: photoUrl,
	}

	photo, err = p.photoRepo.Save(photo)
	return
}

func (p *PhotoSvc) Update(photoInput models.PhotoUpdateInput, photoFileHeader *multipart.FileHeader) (photo models.Photo, err error) {
	// get photo from db to get the photo URL for deletion
	photo, err = p.photoRepo.FindById(int(photoInput.ID))
	if err != nil {
		return
	}

	// if user uploaded a new photo
	if photoFileHeader != nil {
		// get the old photo for deletion
		oldPhotoUrl := photo.PhotoURL

		// validate other input before upload file to cloudinary
		photoInput.PhotoURL = "placeholder"
		_, err = govalidator.ValidateStruct(photoInput)
		if err != nil {
			return
		}

		// open the file and get its content
		photoFile, err := photoFileHeader.Open()
		if err != nil {
			log.Printf("error opening file: %v", err)
			return photo, err
		}
		defer photoFile.Close()

		// upload new photo to cloudinary
		photoUrl, err := helpers.UploadToCloudinary(photoFile)
		if err != nil {
			return photo, err
		}

		// set the photo model for db
		photo = models.Photo{
			Base:     models.Base{ID: photoInput.ID},
			Title:    photoInput.Title,
			Caption:  photoInput.Caption,
			UserID:   photoInput.UserID,
			PhotoURL: photoUrl, // new photo url
		}

		// update data in db
		photo, err = p.photoRepo.Update(photo)
		if err != nil {
			return photo, err
		}

		// delete old photo from cloudinary
		err = helpers.DestroyFromCloudinary(oldPhotoUrl)
		return photo, err
	}

	// if no new photo uploaded, use old photo url & overwrite the other data
	photo = models.Photo{
		Base:     models.Base{ID: photoInput.ID},
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		PhotoURL: photo.PhotoURL, // old photo
		UserID:   photoInput.UserID,
	}

	photo, err = p.photoRepo.Update(photo)
	return
}

func (p *PhotoSvc) Delete(id int) (err error) {
	// get photo from db to get the photo URL for deletion
	photo, err := p.photoRepo.FindById(id)
	if err != nil {
		return
	}

	// delete photo from cloudinary
	err = helpers.DestroyFromCloudinary(photo.PhotoURL)
	if err != nil {
		return
	}

	// delete photo from db
	err = p.photoRepo.Delete(photo)
	return
}
