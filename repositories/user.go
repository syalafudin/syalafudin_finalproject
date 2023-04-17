package repositories

import (
	"github.com/alvinmdj/mygram-api/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Save(user models.User) (models.User, error)
	FindByEmail(user models.User) (models.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Save(user models.User) (models.User, error) {
	err := u.db.Debug().Create(&user).Error
	return user, err
}

func (u *UserRepo) FindByEmail(user models.User) (models.User, error) {
	err := u.db.Debug().Where("email = ?", user.Email).Take(&user).Error
	return user, err
}
