package users

import (
	"errors"
	"sync"

	"github.com/ErmoGarcia/mus-backend/db"
	"github.com/ErmoGarcia/mus-backend/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

var userService *UserService
var initOnce sync.Once

func GetService() *UserService {
	initOnce.Do(initService)
	return userService
}

func initService() {
	db := db.GetDB()
	userService = NewUserService(db)
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

/**
 * User operations
**/

func (dl *UserService) GetUserByID(uid uint) (models.User, error) {

	var user models.User
	var err error

	err = dl.db.First(&user, uid).Error
	if err != nil {
		return user, errors.New("User not found!")
	}

	user.Password = ""
	return user, nil
}

func (dl *UserService) GetUserByName(username string) (models.User, error) {

	var user models.User
	var err error

	err = dl.db.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return user, errors.New("User not found!")
	}

	user.Password = ""
	return user, nil
}

func (dl *UserService) CreateUser(user *models.User) error {

	var err error

	err = dl.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}
